use anyhow::Result;
use log::{error, info};
use rppal::gpio::Gpio;
use std::time::Duration;
use tokio::time;
use std::process::Command;

#[tokio::main]
async fn main() -> Result<()> {
    // Initialize logging
    env_logger::init();
    
    // If your primary network interface isn't eth0, substitute appropriately:
    let mac_address = std::fs::read_to_string("/sys/class/net/eth0/address")?
        .trim()
        .to_string();
    
    // Create an HTTP client
    let client = reqwest::Client::new();
    
    // Initialize GPIO
    let gpio = Gpio::new()?;
    // Example: Using GPIO18 as input (adjust pin number as needed)
    let pin = gpio.get(18)?.into_input();
    
    info!("Controller started, monitoring GPIO...");

    loop {
        // Read sensor data
        let value = pin.is_high();
        
        // Send ping request with sensor data
        match client
            .post("http://your-api-endpoint/ping")
            .json(&serde_json::json!({
                "sensor_value": value,
                "timestamp": chrono::Utc::now().to_rfc3339(),
                "mac_address": mac_address,
            }))
            .send()
            .await
        {
            Ok(response) => {
                if response.status().is_success() {
                    info!("Ping sent successfully");
                } else {
                    error!("Ping failed with status: {}", response.status());
                }
            }
            Err(e) => error!("Failed to send ping: {}", e),
        }

        // Wait before next reading
        time::sleep(Duration::from_secs(5)).await;
    }

    // After your existing loop or at the appropriate place in your code:
    // (Below is an example that we place after some initialization logic.)
    let health_file_path = "/tmp/controller.health"; 
    let mut healthy_detected = false;
    let poll_interval = Duration::from_secs(5);

    for attempt in 1.. {
        if let Ok(contents) = std::fs::read_to_string(health_file_path) {
            if contents.trim() == "healthy" {
                info!("Detected healthy controller on attempt #{}", attempt);
                healthy_detected = true;
                break;
            }
        }
        info!("Controller not yet healthy, attempt #{}...", attempt);
        time::sleep(poll_interval).await;
    }

    if healthy_detected {
        info!("Controller is healthy. Proceeding with post-health logic...");
        // TODO: Add logic here to "destroy old code," handle upgrades, etc.
    } else {
        error!("Failed to detect healthy controller. Exiting...");
        // Handle failure scenario if needed
    }

    // Wait until the old controller is safe to take down, or the system is ready for upgrade.
    // For example, you might check an "idle" file or just proceed immediately after a certain condition.
    // (In a real scenario, you'd put your own logic here.)
    
    // Once ready, attempt the upgrade:
    if let Err(e) = perform_upgrade() {
        error!("Upgrade process encountered an error: {:?}", e);
    }

    Ok(())
}

fn perform_upgrade() -> Result<()> {
    // Paths to old, new, and backup
    let old_path = "/usr/local/bin/controller";
    let backup_path = "/usr/local/bin/controller-backup";
    let new_path = "/tmp/new_controller.bin"; // assume it was downloaded here

    info!("Starting upgrade...");

    // Step 1: backup the old version
    if std::path::Path::new(old_path).exists() {
        info!("Backing up old controller from {} to {}", old_path, backup_path);
        std::fs::rename(old_path, backup_path)?;
    }

    // Step 2: put the new version in place
    if std::path::Path::new(new_path).exists() {
        info!("Moving new controller from {} to {}", new_path, old_path);
        std::fs::rename(new_path, old_path)?;
        // Make sure it's executable
        let mut perms = std::fs::metadata(old_path)?.permissions();
        perms.set_mode(0o755);
        std::fs::set_permissions(old_path, perms)?;
    } else {
        error!("No new controller binary found at {}", new_path);
        return Ok(()); // or return an error
    }

    // Step 3: start the new controller
    info!("Launching the new controller...");
    let mut child = Command::new(old_path)
        .spawn()
        .map_err(|e| anyhow::anyhow!("Failed to start new controller: {}", e))?;

    // Step 4: wait briefly, do a quick health check
    std::thread::sleep(std::time::Duration::from_secs(3));
    if !check_new_controller_health() {
        error!("New controller is not healthy. Rolling back...");
        // kill new process
        let _ = child.kill();
        // revert
        rollback(old_path, backup_path)?;
        return Err(anyhow::anyhow!("New controller did not pass health check"));
    }

    info!("New controller appears healthy. Removing old backup...");
    // If we trust the new version now, remove old backup
    if std::path::Path::new(backup_path).exists() {
        std::fs::remove_file(backup_path)?;
    }

    info!("Upgrade completed successfully!");
    Ok(())
}

/// Check if the newly started controller is healthy enough. In reality, this might
/// call an HTTP endpoint or read from your /tmp/controller.health file, etc.
fn check_new_controller_health() -> bool {
    // Example: pretend we read from /tmp/controller.health
    if let Ok(contents) = std::fs::read_to_string("/tmp/controller.health") {
        return contents.trim() == "healthy";
    }
    false
}

/// Revert to the old binary if the new one failed to start or pass checks.
fn rollback(old_path: &str, backup_path: &str) -> Result<()> {
    if std::path::Path::new(backup_path).exists() {
        info!("Restoring old controller from backup...");
        // remove any partially-upgraded controller
        let _ = std::fs::remove_file(old_path);
        // rename backup -> controller
        std::fs::rename(backup_path, old_path)?;
        // optional: reset permissions
        let mut perms = std::fs::metadata(old_path)?.permissions();
        perms.set_mode(0o755);
        std::fs::set_permissions(old_path, perms)?;
    }
    Ok(())
}

