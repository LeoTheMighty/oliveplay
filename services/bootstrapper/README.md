# Bootstrapper

This is the bootstrapper service. It is responsible for booting up the system and starting the controller service. Also handles updating the system and retrying the controller service until it is in a healthy state.

## How to run

```bash
cargo run
```

## Upgrading the Controller

1. The new controller binary should be placed at `/tmp/new_controller.bin`.  
2. The bootstrapper will rename the existing `/usr/local/bin/controller` to `/usr/local/bin/controller-backup`, then move the new file into `/usr/local/bin/controller`.  
3. If the new controller proves healthy, the bootstrapper will remove the backup; otherwise it will revert.  