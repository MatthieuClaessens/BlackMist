<div align="center">
  <div style="background: linear-gradient(135deg, #1a1a1a 0%, #333333 100%); padding: 50px 20px; border-radius: 15px; margin-bottom: 20px;">
    <h1 style="color: white; margin: 0; font-size: 4em; font-family: sans-serif; letter-spacing: 4px; text-shadow: 2px 2px 4px rgba(0,0,0,0.5);">BlackMist</h1>
    <p style="color: rgba(255,255,255,0.9); margin-top: 10px; font-size: 1.4em; font-weight: 300;">Advanced System-Wide Anonymity Node</p>
  </div>

  <h3>🛡️ Cybersecurity & Software Engineering Project | 2026</h3>
  
  <p>
    <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go" /> &nbsp;
    <img src="https://img.shields.io/badge/React-2026-61DAFB?style=for-the-badge&logo=react" /> &nbsp;
    <img src="https://img.shields.io/badge/Wails-v2-FF4C54?style=for-the-badge&logo=wails" /> &nbsp;
    <img src="https://img.shields.io/badge/Tor_Network-Implementation-7D4698?style=for-the-badge&logo=torproject" />
    <img src="https://img.shields.io/badge/Windows_Kernel-Registry-0078D6?style=for-the-badge&logo=windows" />
  </p>
</div>

---

## 📖 Overview
**BlackMist** is a professional-grade privacy tool designed to bridge the gap between high-level anonymity networks and OS-level networking. Unlike standard browser-based solutions, BlackMist orchestrates the **Tor Expert Bundle** to enforce a global proxy state across the entire Windows environment.

> **💡 Engineering Decision:** The core is built with **Go** to leverage low-level system calls (Win32 API) for registry manipulation, ensuring that the proxy toggle is atomic and leaves no room for IP leaks during state transitions.

---

## 📸 Technical Showcase (Frontend & State Management)

The interface, prototyped on **Figma**, focuses on "One-Click Security". It provides a minimalist aesthetic while maintaining real-time technical feedback (Bootstrap progress, Circuit latency). The core challenge was synchronizing the backend Tor process state with the frontend UI.

### 🌐 System Proxy Toggle & Real-time IP Monitoring

| 🔴 **Mode: Offline (Exposed)** | 🔵 **Mode: Online (Secured)** |
| :--- | :--- |
| <a href="https://i.ibb.co/Qv0pgkPz/image.png" target="_blank"><img src="https://i.ibb.co/Qv0pgkPz/image.png" alt="BlackMist exposed mode"></a> | <a href="https://ibb.co/XkvMhpqf" target="_blank"><img src="https://i.ibb.co/TqFdVRsD/image.png" alt="BlackMist secured mode"></a> |
| *The proxy is disabled. The system IP is exposed (Status: Red).* | *Tor is bootstrapped. System traffic is routed (Status: Blue).* |

> **Pro Tip:** I implemented a **Bootstrap Guard** in Go. The system proxy is only engaged once the Tor circuit reaches **100% connectivity**, preventing any data leaks (DNS or Traffic) during the handshake phase.

---

## ⚙️ 1. Backend & System Integration

BlackMist interacts directly with the **OS Layer** to ensure total traffic encapsulation:
1. **Process Orchestration**: Secure spawning of the Tor child process with detached handles.
2. **Dynamic Configuration**: Generation of volatile configuration files to ensure a **Zero-Log** footprint.
3. **Signal Communication**: Implementation of the Tor Control Protocol to allow real-time identity rotation (NEWNYM).

### 🛠️ Tech Stack
| Layer | Technologies |
| :--- | :--- |
| **Backend Core** | **Go (Golang) 1.21+** |
| **Frontend Framework** | **React / TypeScript / Tailwind CSS** |
| **Native Bridge** | **Wails v2** (Context binding & Events) |
| **Security Layer** | **Tor Expert Bundle**, Windows Registry API |
