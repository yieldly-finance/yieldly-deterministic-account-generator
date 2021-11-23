import utils.wallet as wallet
import utils.accounts as account

def init():   
    print('\n****************************************************')
    print('Welcome to Yieldlys Deterministic Account Generator!')
    print('****************************************************\n')
    print('Please choose carefully:')
    print('1) Mnemonic-to-derivation key')
    print('2) Derivation key-to-Mnemonic key')
    print('3) Generate Accounts')
    print('4) Get Account Mnemonic')

    x = input("\nChoose your option: ")
      
    if(int(x) == 1):
        key = input("\nEnter your Mnemonic: ")
        print("\nYour derivation master key: " + wallet.master_to_derivation(key))
    if(int(x) == 2):
        key = input("\nEnter your derivation key: ")
        print("\nYour mneumonic: " + wallet.derivation_to_master(key))
    if(int(x) == 3):
        seed = input("\nEnter your seed (mnemonic or derivation key): ")
        index = input("\nEnter the index of the account list (max 2^63-1): ")
        amount = input("\nEnter the amount of addresses you would like after the index (max 2^63-1 - index): ")
        print('\nYour accounts: \n')
        print(*account.generateAccountList(seed, int(index), int(amount)), sep = "\n")
    if(int(x) == 4):
        seed = input("\nEnter your seed (mnemonic or derivation key): ")
        index = input("\nEnter the index of the account list (max 2^63-1): ")
        print('\nYour accounts mnemonic: \n')
        print(*account.generateAccountMneumonic(seed, int(index)), sep = "\n")

init()
