# 🎨 Assignment Showcase: Art Protection, Trading & Gaming Platform

This project demonstrates how we apply what we’ve learned to build an **art protection, trading, and gaming platform** using Hyperledger Fabric.

## ✨ Features
- **Artist Mode**:  
  Upload your works, which will be securely protected on the platform.
- **Card Minting**:  
  Players use platform-issued currency to **randomly mint cards**, where the card faces display the artists’ works.
- **Unique Attributes**:  
  Each card comes with **random attributes and unique effects**.
- **Trading & Strategy**:  
  Players can trade cards to form factions, enhancing offensive and defensive strategies.
- **Fabric-Powered Trust**:  
  All card and transaction information is **traceable and authentic** via Fabric.

---

## 🚀 Getting Started

### Start the Fabric Network
```bash
cd /your_fabric_path/fabric-samples/test-network

# Clean up old network
./network.sh down

# Start a new channel with CA
./network.sh up createChannel -c mychannel -ca

# Deploy chaincode
./network.sh deployCC -ccn basic -ccp /your_file_path/AssignmentOfFabric/chaincode -ccl go

```
---

✅ Now your backend is running and connected to Fabric!  
⚠️ **Don’t forget to update the file paths in `fabric.go` to match your environment.**
