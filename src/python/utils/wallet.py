from algosdk import mnemonic

def master_to_derivation(_mnemonic: str):
    return mnemonic.to_master_derivation_key(_mnemonic)

def derivation_to_master(_key: str):
    return mnemonic.from_master_derivation_key(_key)