# yieldly-deterministic-account-generator

If you wish to test the compatibility with kmd, please refer to one of the following options to install algod + kmd to verify that the accounts you are generating are indeed deterministic.

Sandbox (RECOMMENDED): [https://github.com/algorand/sandbox](https://github.com/algorand/sandbox)

Private Network: [https://developer.algorand.org/tutorials/create-private-network](https://developer.algorand.org/tutorials/create-private-network)

Pull the code from the repository and open a terminal in the projects workspace. Examples for each programming language have been provided. Under the /src folder of the repository sits each respective programming languages folder. Due to the deliberate similarities of the software stack (every function has a mirrored version in terms of functionality on each programming language), choosing ultimately comes down to your preferred programming language. 

You then proceed to /cd into your preferred languages folder and install all relevant packages.
Run the program file sitting under the /prog_language folder to get started. For example, with Python, this will be:

```powershell
python3 main.py
```

You should then see a screen similar to the image below:

[screenshot 1](./docs/1.png)

### GENERATING MNEMONIC FROM DERIVATION MASTER KEY

Generate a new wallet like so (in this example sandbox will be used):

[screenshot 2](./docs/2.png)

Select yes to being able to see your wallets mnemonic, this will generate a 25 word sentence that you will need to note down somewhere. After obtaining your wallets mnemonic, run the program as described above and type "1" into the option selection.

[screenshot 3](./docs/3.png)

You will then be asked to input your mnemonic phrase. Once you have entered all 25 words, press enter again, you should see your master key generated in base64 format in the output of the terminal.

[screenshot 4](./docs/4.png)

*It is not recommended that you share this key or your mnemonic around. If anyone gets their hands on either your mnemonic or the derived master key, then they will be able to recreate the mnemonics / private keys of any account derived from the methods shown later in this document.*

### GENERATING DERIVATION MASTER KEY FROM MNEMONIC

Select option 2:

[screenshot 5](./docs/5.png)

Either by using the output generated from the above method or a key that you have managed to derive from another way, input your master key in base64 format. The console should then output your wallets mnemonic:

[screenshot 6](./docs/6.png)

If you have used the method described under **"GENERATING MNEMONIC FROM DERIVATION MASTER KEY"** you should notice that the mnemonic that you entered is the same that was output in the console from this method. Showcasing that the system performs as expected and means you can freely call between the two functions.

*It is not recommended that you share this key or your mnemonic around. If anyone gets their hands on either your mnemonic or the derived master key, then they will be able to recreate the mnemonics / private keys of any account derived from the methods shown later.*

### GENERATING ACCOUNT ADDRESSES FROM EITHER DERIVATION MASTER KEY OR MNEMONIC
Select option 3:

[screenshot 7](./docs/7.png)

You will then be asked to enter your mnemonic phrase or master key, followed by the index of where you would like to start (for a deeper look into how the accounts are generated, please go to the "How it works" section), for testing purposes, it is recommended that you start at index 1. Lastly you will be prompted on how many accounts after you would like to generate after this index. Please note that the maximum number of accounts that can be generated is 2⁶³ - 1.

If all your information is correct and in range, it should print a list of addresses as shown below:

[screenshot 8](./docs/8.png)

If you have this wallet imported in your system, you should be able to generate the same addresses to verify that the program works.

### GENERATING ACCOUNT MNEMONICS FROM EITHER DERIVATION MASTER KEY OR MNEMONIC
Select option 4. Similar to the above option, you will then be asked to enter your mnemonic phrase or master key, followed by the index of where you would like to start (for more information, please go to the "How it works" section), for testing purposes, it is recommended that you start at index 1. Unlike generating account private keys, this example has been limited to producing 1 at a time. This is primarily for safety reasons for new users just wanting to learn how the system works without accidently exposing every mnemonic in their wallet.

If all your information is correct and in range, it should print a list of addresses as shown below:

[screenshot 9](./docs/9.png)

If you have this wallet imported in your system, you should be able to export the mnemonic to verify that both are equal and thus, verifies that the system can indeed generate accounts deterministically that are ready for use in the Algorand ecosystem.


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)