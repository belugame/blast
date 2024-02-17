# BLAST

<img src="logo.png" width=256px height=256px alt="Blast logo" title="Blast logo">

## Cast your Linux audio to DLNA receivers

You need `pactl`, `parec` and `ffmpeg` executables/dependencies on your system to run Blast.

If you have all that then you can launch `blast` and it looks like this when you run it:

```
[user@user blast]$ ./blast 
----------
DLNA receivers
0: Kitchen
1: Phone
2: Bedroom
3: Livingroom TV
----------
Select the DLNA device:
[1]
----------
Audio sources
0: alsa_output.pci-0000_00_1b.0.analog-stereo.monitor
1: alsa_input.pci-0000_00_1b.0.analog-stereo
2: bluez_output.D8_AA_59_95_96_B7.1.monitor
3: blast.monitor
----------
Select the audio source:
[2]
----------
Your LAN ip addresses
0: 192.168.1.14
1: 192.168.122.1
2: 2a04:ec00:b9ab:555:3c50:e6e8:8ea:211f
3: 2a04:ec00:b9ab:555:806d:800b:1138:8b1b
4: fe80::f4c2:c827:a865:35e5
----------
Select the lan IP address for the stream:
[0]
----------
2023/07/08 23:53:07 starting the stream on port 9000 (configure your firewall if necessary)
2023/07/10 23:53:07 stream URI: http://192.168.1.14:9000/stream.mp3
2023/07/08 23:53:07 setting av1transport URI and playing
```

There's also a `-debug` and `-headers` flags if you want to inspect your DLNA device. Also `-log` to inspect what parec and ffmpeg is doing. 

### Non-interactive usage and extra flags

```
[ugjka@ugjka blast]$ blast -h
Usage of blast:
  -bige
        use big endian for capture and raw formats
  -bitrate int
        audio format bitrate (default 320)
  -bits int
        audio bitdepth (default 16)
  -channels int
        audio channels (default 2)
  -chunk int
        chunked size in seconds (default 1)
  -debug
        print debug info
  -device string
        dlna device's friendly name
  -dummy
        only serve content
  -format string
        stream audio format (default "mp3")
  -headers
        print request headers
  -ip string
        host ip address
  -log
        log parec and ffmpeg stderr
  -mime string
        stream mime type (default "audio/mpeg")
  -nochunked
        disable chunked tranfer endcoding
  -port int
        stream port (default 9000)
  -rate int
        audio sample rate (default 44100)
  -source string
        audio source (pactl list sources short | cut -f2)
  -useaac
        use aac audio
  -useflac
        use flac audio
  -uselpcm
        use lpcm audio
  -usewav
        use wav audio
```

## Building

You need the `go` and `go-tools` toolchain, also `git`

then execute:

```
git clone https://github.com/ugjka/blast
cd blast
go build
```

now you can run blast with:
```
[user@user blast]$ ./blast
```

## Bins

Prebuilt Linux binaries are available on the releases [page](https://github.com/ugjka/blast/releases)

## Why not use pulseaudio-dlna?

This is for pipewire-pulse users.

## Caveats

* You need to allow port 9000 from LAN for the DLNA receiver to be able to access the HTTP stream, you can change it with `-port` flag
* blast monitor sink may not be visible in the pulse control applet unless you enable virtual streams

## Trivia

What on earth is "x-rincon-mp3radio"

## License

```
MIT+NoAI License

Copyright (c) 2023 ugjka <ugjka@proton.me>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights/
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

This code may not be used to train artificial intelligence computer models
or retrieved by artificial intelligence software or hardware.
```
