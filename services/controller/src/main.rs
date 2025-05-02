use anyhow::Result;
use log::{error, info};
use rppal::gpio::Gpio;
use std::time::Duration;
use tokio::time;

#[tokio::main]
async fn main() -> Result<()> {
    // Initialize logging
    env_logger::init();
    
    // If your primary network interface isn't eth0, substitute appropriately:
    let mac_address = std::fs::read_to_string("/sys/class/net/wlan0/address")?
        .trim()
        .to_string();

    // Create an HTTP client
    let client = reqwest::Client::new();
    
    let gpio = Gpio::new()?;
    let mut input_pins: Vec<(u8, InputPin)> = Vec::new();

    for pin_id in 0u8..=27 {
        match gpio.get(pin_id) {
            Ok(pin) => input_pins.push((pin_id, pin.into_input())),
            Err(e)  => warn!("Skipping GPIO{pin_id} – {e}"),
        }
    }

    info!(
        "Controller started, monitoring {} GPIO pins...",
        input_pins.len()
    );

    loop {
        // ────────────────────────────────────────────────────────────────
        // 2. Read every pin’s level and print it.
        // ────────────────────────────────────────────────────────────────
        let mut pin_states = serde_json::Map::new();

        for (pin_id, pin) in &input_pins {
            let high = pin.is_high();
            info!("GPIO{:02}: {}", pin_id, if high { "High" } else { "Low" });

            // keep it around so we can also send everything in one payload
            pin_states.insert(format!("GPIO{pin_id}"), serde_json::json!(high));
        }
        
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
}

