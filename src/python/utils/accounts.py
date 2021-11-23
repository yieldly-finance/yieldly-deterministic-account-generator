from Cryptodome.Hash import SHA512
import hmac
from ed25519.keys import SigningKey
import base64
import utils.wallet as wallet
from algosdk import mnemonic

buffer = lambda x: x

def hkdf_expand(pseudo_random_key, info=b"", length=32, hash=SHA512.new(truncate="256")):
	'''
	Expand `pseudo_random_key` and `info` into a key of length `bytes` using
	HKDF's expand function based on HMAC with the provided hash (default
	SHA-512). See the HKDF draft RFC and paper for usage notes.
	'''
	hash_len = hash.digest_size
	length = int(length)
	if length > 255 * hash_len:
		raise Exception("Cannot expand to more than 255 * %d = %d bytes using the specified hash function" %\
			(hash_len, 255 * hash_len))
	blocks_needed = length // hash_len + (0 if length % hash_len == 0 else 1) # ceil
	okm = b""
	output_block = b""
	for counter in range(blocks_needed):
		output_block = hmac.new(pseudo_random_key, buffer(output_block + info + bytearray((counter + 1,))),\
			hash).digest()
		okm += output_block
	return okm[:length]

def generateAccountList(seed: str, index: int, amount: int = 1):
    accountArray = []
    decoded_key = ""

    if(seed.__len__() > 50):
        master_key = wallet.master_to_derivation(seed)

        decoded_key = base64.b64decode(master_key)
    else:
        decoded_key = base64.b64decode(seed)


    for x in range(index, index + amount):
        stringTemp = "AlgorandDeterministicKey-%d" % (x)

        #Create the info string
        info = bytes(stringTemp, 'utf-8')

        key = hkdf_expand(decoded_key, info, 32, SHA512.new(truncate="256"))
        #print(key)
        
        sk = SigningKey(key)
        vk = sk.get_verifying_key()

        chksum = SHA512.new(truncate="256")
        chksum.update(vk.to_bytes())

        # Gets the last 4 bytes
        test: bytes = vk.to_bytes() + chksum.digest()[-4:]
        tempString = base64.b32encode(test)

        accountArray.append(tempString[slice(0, 58)].decode("ascii"))

    return accountArray


def generateAccountMneumonic(seed: str, index: int):
    accountArray = []
    decoded_key = ""

    if(seed.__len__() > 50):
        master_key = wallet.master_to_derivation(seed)

        decoded_key = base64.b64decode(master_key)
    else:
        decoded_key = base64.b64decode(seed)

    for x in range(index, index + 1):
        stringTemp = "AlgorandDeterministicKey-%d" % (x)

        #Create the info string
        info = bytes(stringTemp, 'utf-8')

        key = hkdf_expand(decoded_key, info, 32, SHA512.new(truncate="256"))
        
        sk = SigningKey(key)

        tempString2 = base64.b64encode(sk.to_bytes())

        accountArray.append(mnemonic.from_private_key(tempString2))

    return accountArray


