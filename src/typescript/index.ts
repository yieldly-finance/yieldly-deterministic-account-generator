import { generateAccountMnemonic, generateAccountList } from "./utils/account";
import { derivationToMaster, masterToDerivation } from "./utils/wallet";

import readline from "readline";

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

const main = (): void => {
  console.log("\n****************************************************");
  console.log("Welcome to Yieldlys Deterministic Account Generator!");
  console.log("****************************************************\n");
  console.log("Please choose carefully:");
  console.log("1) Mnemonic-to-derivation key");
  console.log("2) Derivation key-to-Mnemonic key");
  console.log("3) Generate Accounts");
  console.log("4) Get Account Mnemonic");

  rl.question("\nChoose your option: ", (i) => {
    switch (i) {
      case "1":
        rl.question("\nEnter your Mnemonic: ", (mnemonic) => {
          console.log(
            "\n" + Buffer.from(masterToDerivation(mnemonic)).toString("base64")
          );
        });
      case "2":
        rl.question("\nEnter your derivation key: ", (key) => {
          console.log("\n" + derivationToMaster(key));
        });
      case "3":
        rl.question(
          "\nEnter your seed (mnemonic or derivation key): ",
          (key) => {
            rl.question(
              "\nEnter the index of the account list (max 2^63-1): ",
              (index) => {
                rl.question(
                  "\nEnter the amount of addresses you would like after the index (max 2^63-1 - index): ",
                  (amount) => {
                    console.log(
                      "\n" +
                        generateAccountList(key, Number(index), Number(amount))
                    );
                  }
                );
              }
            );
          }
        );
      case "4":
        rl.question(
          "\nEnter your seed (mnemonic or derivation key): ",
          (key) => {
            rl.question(
              "\nEnter the index of the account list (max 2^63-1): ",
              (index) => {
                console.log("\n" + generateAccountMnemonic(key, Number(index)));
              }
            );
          }
        );
    }

    return;
  });
};

main();
