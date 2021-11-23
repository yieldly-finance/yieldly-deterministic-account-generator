import algosdk from "algosdk";
import nacl from "tweetnacl";
import { masterToDerivation } from "./wallet";
import hkdf from "futoin-hkdf";
import sha512 from "js-sha512";
import * as utils from "./encoders";
import base32 from "hi-base32";

export const generateAccountList = (
  seed: string,
  index: number,
  amount: number
): string[] => {
  let accountArray: string[] = [];
  let decodeKey: Uint8Array;

  if (seed.length > 50) {
    let masterKey = masterToDerivation(seed);
    decodeKey = masterKey;
  } else {
    decodeKey = new Uint8Array(Buffer.from(seed, "base64"));
  }

  for (var i = index; i < index + amount; i++) {
    let info = "AlgorandDeterministicKey-1";

    let key = hkdf.expand(
      "SHA512-256",
      hkdf.hash_length("sha256"),
      Buffer.from(decodeKey),
      32,
      info
    );

    let { publicKey } = nacl.sign.keyPair.fromSeed(key);

    const checksum = sha512.sha512_256.array(publicKey).slice(28, 32);

    const addr = base32.encode(utils.concatArrays(publicKey, checksum));

    accountArray.push(addr.toString().slice(0, 58));
  }

  return accountArray;
};

export const generateAccountMnemonic = (
  seed: string,
  index: number
): string[] => {
  let accountArray: string[] = [];
  let decodeKey: Uint8Array;

  if (seed.length > 50) {
    let masterKey = masterToDerivation(seed);
    decodeKey = masterKey;
  } else {
    decodeKey = new Uint8Array(Buffer.from(seed, "base64"));
  }

  for (var i = index; i < index + 1; i++) {
    let info = "AlgorandDeterministicKey-1";

    let key = hkdf.expand(
      "SHA512-256",
      hkdf.hash_length("sha256"),
      Buffer.from(decodeKey),
      32,
      info
    );

    let { secretKey } = nacl.sign.keyPair.fromSeed(key);

    accountArray.push(algosdk.secretKeyToMnemonic(secretKey));
  }

  return accountArray;
};
