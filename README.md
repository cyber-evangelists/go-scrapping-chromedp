# Web Scraping with ChromeDP in Go

## Overview
This project demonstrates web scraping using ChromeDP in Go to extract ASN (Autonomous System Number) information from a specific webpage.

## Project Structure
### Files and Directories

- **main.go**: Entry point for the application.
- **new.go**: Alternative approach for scraping .
  

## Installation

### Prerequisites
- Go 1.16+
- Chrome or Chromium browser
- Dependencies managed via `go mod`

### Steps
1. Clone the repository:
    ```sh
    git clone https://github.com/your-repo/web-scraping.git
    cd web-scraping

    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Ensure Chrome or Chromium is installed and accessible.


## Usage

### Start the Servers
This script uses a default ChromeDP setup to scrape ASN information:
```sh
go run main.go
 ```

This script demonstrates an alternate setup with a custom user agent:
```sh
go run new.go
 ```

## Output
Data extracted from the webpage is saved in `as_data.json` in JSON format.

## Customization
- Adjust the webpage URL and CSS selectors in `main.go` and `new.go` as per your scraping needs.
- Modify data parsing logic in `main.go` to extract different types of information.

## Dependencies
**chromedp/chromedp:** Go package for Chrome DevTools Protocol.
**chromedp/cdproto/cdp:** Go implementation of Chrome DevTools Protocol.


## Conclusion

This project showcases the use of ChromeDP in Go for web scraping tasks, providing flexibility and control over browser interactions. By following the provided examples in `main.go` and `new.go`, developers can extract structured data from web pages efficiently. The project demonstrates how to set up ChromeDP, customize scraping logic, and handle extracted data for further processing or storage. Whether scraping for ASN details or other types of information, this setup offers a robust foundation for web scraping projects in Go.

```bash
This app was made with 💖 by Hamza under the guidance of Sir Husnain.
```
