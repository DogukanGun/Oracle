// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;


contract Wallet{

    struct Asset {
        string sign;
        uint256 amount;
    }

    struct UserWallet {
        string passwordSaver;
        string password;
        uint256 assetCount;
        mapping(uint256 => Asset) assets;
    }

    mapping(address => UserWallet) public userWallets;

    function createUserAddress(string memory passwordSaver,string memory password, string memory username) public returns(address){
        address userAddress = generateAddress(username);
        UserWallet storage userWallet = userWallets[userAddress];
        userWallet.password = password;
        userWallet.passwordSaver = passwordSaver;
        userWallet.assetCount = 0;
        return userAddress;
    }

    function getUserAssets(address userAddresss) external view returns (Asset[] memory) {
        UserWallet storage userWallet = userWallets[userAddresss];
        Asset[] memory assets = new Asset[](userWallet.assetCount);
        for (uint256 i = 0; i < userWallet.assetCount; i++) {
            assets[i] = userWallet.assets[i];
        }
        return assets;
    }

    function addNewAsset(address userAddress,string calldata sign) public {
        UserWallet storage userWallet = userWallets[userAddress];
        if(!isAssetExist(userAddress,sign)){
            Asset memory asset;
            asset.sign = sign;
            asset.amount = 0;
            userWallet.assets[userWallet.assetCount] = asset;
            userWallet.assetCount = userWallet.assetCount + 1;
        }
    }

    function depositAsset(address userAddress,string memory sign,uint256 amount) public {
        UserWallet storage userWallet = userWallets[userAddress];
        for (uint256 i = 0; i < userWallet.assetCount; i++) {
            if(compareStrings(userWallet.assets[i].sign,sign)){
                userWallet.assets[i].amount = userWallet.assets[i].amount + amount; 
            }
        }
    }

    function withdrawAsset(address userAddress,string memory sign,uint256 amount) public {
        UserWallet storage userWallet = userWallets[userAddress];
        for (uint256 i = 0; i < userWallet.assetCount; i++) {
            if(compareStrings(userWallet.assets[i].sign,sign)){
                userWallet.assets[i].amount = userWallet.assets[i].amount - amount; 
            }
        }
    }

    function isAssetExist(address userAddress,string memory sign) public  view returns(bool){
        UserWallet storage userWallet = userWallets[userAddress];
        for (uint256 i = 0; i < userWallet.assetCount; i++) {
            if(compareStrings(userWallet.assets[i].sign,sign)){
                return true;
            }
        }
        return false;
    }

    function compareStrings(string memory str1, string memory str2) public pure returns (bool) {
        return keccak256(abi.encodePacked(str1)) == keccak256(abi.encodePacked(str2));
    }

    function generateAddress(string memory username) private pure returns (address) {
        bytes memory usernameBytes = bytes(username);
        address generatedAddress = address(uint160(uint256(keccak256(usernameBytes))));
        return generatedAddress;
    }


}