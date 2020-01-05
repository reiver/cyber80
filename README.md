# cyber80

The **cyber80** is a **fantasy console computer** geared towards **video games** that is inspired by arcade game machines, home video game consoles, handheld game consoles, and (other) computers from the **1980s**.

I.e., although the **cyber80** is similar to machines from the 1980s, it did _not_ actually exist in the 1980s.
It is a _fantasy_ 🕹️

The **cyber80** was designed for:
* teaching programming through game development, and
* retro game development.

## 80

The “80” in “cyber80” stands for “1980s”.
Which is the era that the _cyber80_ is inspired by.

## Video Game Consoles

A **video game console** is a _computer_ designed primarily to be used to play video games.

Early examples of _video game consoles_ include:
* Atari 2600,
* Intellivision,
* ColecoVision,
* Vectrex,
* etc.

Later example of _video game consoles_ included:
* Nintendo Entertainment System (NES),
* Super Nintendo Entertainment System (SNES),
* Sega Master System,
* Sega Genesis,
* Sega Saturn,
* Sega Dreamcast,
* Neo Geo,
* etc.

Later examples of handheld video game consoles include:
* Game Boy,
* Atari Lynx,
* Game Gear,
* etc.

Current examples of _video game consoles_ includ:
* XBox,
* Playstation,
* Nintendo Switch,
* etc.

## Fantasy Video Game Consoles

A **fantasy video game console** is a _make-believe_ _video game console_ that never actually existed,
but is inspired ‘real’ _video game consoles_.

A **fantasy video game console computer** is similar to a **retro game console emulator**, but for a machine that never existed.

## cyber80 Specification

**Display**:
* **Raster Mode**: 128×192 pixels, 16 color palette
* **Vector Mode**: 256×384 points, 16 color palette

**Input**:
* 2 controllers each with 8-buttons

## cyber80 RAM Layout

| Address | Description    | Size         | Note                                                                          |
|---------|----------------|--------------|-------------------------------------------------------------------------------|
| 0x0000  | Raster Display | 24,576 bytes | = 128×192 pixels with 1 byte per pixel; only lower 4 bits of each byte matter |
| 0xC000  | Palette        | 48 bytes     | = 3 bytes per color with 16 colors                                            |
