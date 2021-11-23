import algosdk from "algosdk";
import btoa from "btoa";
import atob from "atob";

export const masterToDerivation = (_mnemonic: string): Uint8Array => {
  return algosdk.mnemonicToMasterDerivationKey(_mnemonic);
};

export const derivationToMaster = (key: string): string => {
  return algosdk.masterDerivationKeyToMnemonic(
    new Uint8Array(Buffer.from(key, "base64"))
  );
};
