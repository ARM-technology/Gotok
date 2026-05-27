# 📱 gotok

A lightweight, high-performance, and **zero-dependency** TikTok automation tool written in **Go** using native **ADB (Android Debug Bridge)**. 

No complex libraries, no background apps, and no external keyboards required. It uses **Dynamic Screen Scaling (Percentage-based Tapping)** to control TikTok flawlessly across any Android device screen resolution.

---

## ✨ Features
* **Zero Dependencies:** Pure Go code utilizing native `os/exec` to talk directly with ADB.
* **Smart Dynamic Scaling:** Never hardcode $X$ and $Y$ coordinates again! Tap points are calculated using relative percentages (`ratioX`, `ratioY`).
* **Safe Text Injection:** Bypasses Android OS clipboard restrictions and types smoothly with full spacebar (` `) support.
* **Pure Automation:** Auto-scroll through videos, auto-open comments, inject text, and interact seamlessly.

---

## 🛠️ Requirements
1. **Go SDK** installed on your PC.
2. **ADB** installed and added to your system Environment Variables (PATH).
3. Android Device with **USB Debugging** enabled via Developer Options.

---

## 🚀 Quick Start

