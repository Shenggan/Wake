# Wake: A Wake on LAN Tool

You can download the binary file for [linux](https://github.com/Shenggan/Wake/releases/download/0.1.0/wake_linux.zip), [windows](https://github.com/Shenggan/Wake/releases/download/0.1.0/wake_win.zip) and [macos](https://github.com/Shenggan/Wake/releases/download/0.1.0/wake_darwin.zip). Or you can build it from source code.

```shell
go build wake.go
```

### Usage

You can just run it directly on linux and macos. On windows, you can click the wake.exe.

```shell
./wake
```

You can use `./wake -h` for more details. And browser will pop up with a page like below:

![](static/example.png)

Accroding to your target machine MAC address and IP address, you can wake it on LAN.

### Acknowledge
Thanks for [mpolden/wakeup](https://github.com/mpolden/wakeup) and [sabhiram/go-wol](https://github.com/sabhiram/go-wol).
