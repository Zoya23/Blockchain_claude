# Practical 10 — Combining Blockchain with IoT for Secure Data Management

## Aim
Simulate IoT sensor readings (temperature, humidity), store them in a hash-linked blockchain, and detect tampering if a reading is altered after the fact.

## ⚠️ Note on this file
Like Practical 9, the journal describes this script's required structure (Block class, Blockchain class, `display_chain()`, `verify_chain()`, random sensor generation, a tampering test) step-by-step, but the actual Python source was never pasted in. [`iot_blockchain.py`](iot_blockchain.py) below is written to match that description and reproduces the exact tampering-test flow shown in your journal (`iot_chain.chain[1].sensor_data = "Temperature=100°C"`).

## Run
```bash
python iot_blockchain.py
```

## What it does
1. `Block` stores `index`, `sensor_data`, `timestamp`, `previous_hash`, and a SHA-256 `hash` of all of those combined
2. `Blockchain` creates a genesis block, then `add_block()` appends new blocks chained to the previous block's hash
3. `generate_sensor_reading()` produces a random temperature (20–40°C) and humidity (30–90%) reading
4. Three sensor readings are added and the chain is displayed and verified
5. **Tampering test**: `iot_chain.chain[1].sensor_data` is directly overwritten to `"Temperature=100°C"`, then `verify_chain()` is run again

### Verified Output
```
IoT Blockchain Verified Successfully

Warning! IoT Data Tampered
```
✅ Ran end-to-end — first verification passes, then the tampering test correctly fails after the direct mutation, exactly matching the journal's described flow.
