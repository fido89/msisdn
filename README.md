# MSISDN PARSER
REST API that takes MSISDN as an input and parsed it to return MNO identifier, country dialling code, subscriber number and country identifier as defined with ISO 3166-1-alpha-2. 

## Instructions
1. Clone this repository
    ```
    git clone https://github.com/fido89/msisdn.git
    ```
2. Start server with
    ```
    make up
    ```
3. Use it
    ```
    http://localhost:8080/msisdn/parse?msisdn=<MSISDN>
    ```

## Data sources
https://github.com/xxxdepy/simple-country-dial-codes
https://github.com/google/libphonenumber/tree/master/resources/carrier/en
