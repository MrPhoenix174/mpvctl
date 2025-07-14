#!/usr/bin/env bash

# === MPVCTL INIT SCRIPT ===
# Installing dependencies and other...

GREEN="\e[32m"
RED="\e[31m"
YELLOW="\e[33m"
RESET="\e[0m"

echo -e "${GREEN}== MPVCTL Initialization Started ==${RESET}"
echo "[i] This script will install required dependencies and build mpvctl."
echo "[i] It will ask for your password to install packages using 'sudo'."

# === Add SUDO ===
if ! sudo -v; then
    echo -e "${RED}❌ Failed to get sudo permissions. Exiting.${RESET}"
    exit 1
fi


install_deps() {
    echo "[*] Detecting your Linux distribution..."

    if [[ -f /etc/os-release ]]; then
        . /etc/os-release
        DISTRO=$ID
    else
        echo -e "${RED}Cannot detect distro. Exiting.${RESET}"
        exit 1
    fi

    echo "[*] Detected: $DISTRO"

    echo "[*] Installing required packages: Go, MPV..."

    case "$DISTRO" in
        arch|manjaro)
            sudo pacman -Sy --noconfirm go mpv
            ;;
        debian|ubuntu)
            sudo apt update && sudo apt install -y golang mpv
            ;;
        fedora)
            sudo dnf install -y golang mpv
            ;;
        void)
            sudo xbps-install -Sy go mpv
            ;;
        *)
            echo -e "${RED}Unsupported distro: $DISTRO${RESET}"
            echo "Please install Go and MPV manually."
            return
            ;;
    esac

    # === Install Cava ===
    read -p "Do you want to install Cava? (Y/n): " install_cava
    install_cava=${install_cava:-Y}
    if [[ "$install_cava" =~ ^[Yy]$ ]]; then
        echo "[*] Installing Cava..."
        case "$DISTRO" in
            arch|manjaro)
                sudo pacman -Sy --noconfirm cava
                ;;
            debian|ubuntu)
                sudo apt install -y cava
                ;;
            fedora)
                sudo dnf install -y cava
                ;;
            void)
                sudo xbps-install -Sy cava
                ;;
            *)
                echo "Please install cava manually."
                ;;
        esac
    fi

    # === Install EasyEffects ===
    read -p "Do you want to install EasyEffects? (Y/n): " install_effects
    install_effects=${install_effects:-Y}
    if [[ "$install_effects" =~ ^[Yy]$ ]]; then
        echo "[*] Installing EasyEffects..."
        case "$DISTRO" in
            arch|manjaro)
                sudo pacman -Sy --noconfirm easyeffects
                ;;
            debian|ubuntu)
                sudo apt install -y easyeffects
                ;;
            fedora)
                sudo dnf install -y easyeffects
                ;;
            void)
                sudo xbps-install -Sy easyeffects
                ;;
            *)
                echo "Please install EasyEffects manually."
                ;;
        esac
    fi
}

# === Call Installer 
install_deps

# === Config === 
echo "[*] Creating config files..."

mkdir -p config

LINKS_FILE="config/links.json"
SETTINGS_FILE="config/settings.json"

if [[ ! -f "$LINKS_FILE" ]]; then
    echo "[]" > "$LINKS_FILE"
    echo -e "${GREEN}Created $LINKS_FILE${RESET}"
fi

if [[ ! -f "$SETTINGS_FILE" ]]; then
    cat <<EOF > "$SETTINGS_FILE"
{
  "autoplay": true,
  "startEasyEffects": false,
  "startCava": false
}
EOF
    echo -e "${GREEN}Created $SETTINGS_FILE${RESET}"
fi

# === Build === 
echo "[*] Building project..."
go build -o mpvctl

# === Add to PATH === 
add_to_path=false
read -p "Do you want to add 'mpvctl' to PATH (/usr/local/bin)? (Y/n): " response
response=${response:-Y}
if [[ "$response" =~ ^[Yy]$ ]]; then
    sudo cp mpvctl /usr/local/bin/
    echo -e "${GREEN}Added 'mpvctl' to /usr/local/bin${RESET}"
    add_to_path=true
else
    echo -e "${YELLOW}You can still use mpvctl by running it from the current folder.${RESET}"
    echo -e "${YELLOW}To run it from anywhere later, re-run this script and agree to add it to PATH.${RESET}"
fi

# === Finish === 
echo
echo -e "${GREEN}== Installation Complete ==${RESET}"

if $add_to_path; then
    echo -e "${GREEN}You can now run 'mpvctl' from any terminal.${RESET}"
else
    echo -e "${YELLOW}Run './mpvctl' from the current folder to start.${RESET}"
fi

echo -e "${YELLOW}Note: This project is in development, features may change.${RESET}"
