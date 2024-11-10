# Ascii-Art-Justify

This program transforms user-input text into customizable ASCII art, allowing for various alignment options and banner styles. The program adapts the output to fit terminal sizes, ensuring responsive and visually appealing representations.

## Documentation

This section illustrates how to make use of this program.

### Installation

To run this program, download and install the latest version of Go from [here](https://go.dev/doc/install).

### Usage

1. Clone this repository onto your terminal by using the following command:
    ```bash
    git clone https://learn.zone01kisumu.ke/git/kada/ascii-art-justify
    ```

2. Navigate into the ascii-art-justify directory by using the command:
    ```bash
    cd ascii-art-justify
    ```

3. To run the program, execute the command below, specifying the alignment option and the string:
    ```bash
    go run . --align=<type> "your text" <banner>
    ```
   Replace `<type>` with one of the following alignment options: `center`, `left`, `right`, or `justify`. Replace `<banner>` with the desired banner format.

   Example:
   ```bash
   go run . --align=center "hello" standard
   ```
   Expected output:



                              _              _   _          
                             | |            | | | |         
                             | |__     ___  | | | |   ___   
                             |  _ \   / _ \ | | | |  / _ \  
                             | | | | |  __/ | | | | | (_) | 
                             |_| |_|  \___| |_| |_|  \___/  

### Features
- --align=center: Center aligns the ASCII art.
- --align=left: Left aligns the ASCII art.
- --align=right: Right aligns the ASCII art.
- --align=justify: Justifies the ASCII art.

### Contributions

Pull requests are welcome! Users of this program are encouraged to contribute by adding features or fixing bugs. For major changes, please open an issue first to discuss your ideas.

### Authors
[josotieno](https://learn.zone01kisumu.ke/git/josotieno/)

[kada](https://learn.zone01kisumu.ke/git/kada/)

## Licence
[MIT License](./LICENSE)
## Credits
[Zone01Kisumu](https://www.zone01kisumu.ke/)