### Installing chrome in debian based linux

- RUN `apt-get update && apt-get upgrade -y && apt-get install -y  ca-certificates`
- RUN `wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb && apt install -y ./google-chrome-stable_current_amd64.deb`

- ENV `CHROME_BIN=/usr/bin/google-chrome \
    CHROME_PATH=/usr/bin/google-chrome


### Installing chrome in fedora based linux

- RUN `yum update && yum upgrade -y && yum install -y ca-certificates`

- RUN `wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm && yum localinstall -y ./google-chrome-stable_current_x86_64.rpm`

- ENV CHROME_BIN=/usr/bin/google-chrome \
    CHROME_PATH=/usr/bin/google-chrome


### Install chrome in amazon linux 2023
- RUN `dnf update -y && dnf upgrade -y && dnf -y install wget gtk3-devel liberation-fonts nss-tools vulkan-loader xdg-utils alsa-lib-devel`

- RUN `wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm && rpm -i ./google-chrome-stable_current_x86_64.rpm`
- ENV CHROME_BIN=/usr/bin/google-chrome \
    CHROME_PATH=/usr/bin/google-chrome

