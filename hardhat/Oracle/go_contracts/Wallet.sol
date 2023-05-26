// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;


contract Wallet{

    struct AssetPerChain {
        uint128 sign;
        uint128 chainId;
        uint256 amount;
    }

    struct Asset{
        uint128 sign;
        uint256 amount;
    }

    struct UserWallet {
        uint256 assetCount;
        uint256 assetCountWithoutChain;
        mapping(uint256 => AssetPerChain) assets;
        mapping(uint256 => Asset) assetsWithoutChain;

    }

    mapping(address => UserWallet) public userWallets;

    function getUserAssets(address userAddresss) external view returns (AssetPerChain[] memory) {
        UserWallet storage userWallet = userWallets[userAddresss];
        AssetPerChain[] memory assets = new AssetPerChain[](userWallet.assetCount);
        for (uint256 i = 0; i < userWallet.assetCount; i++) {
            assets[i] = userWallet.assets[i];
        }
        return assets;
    }

    function getUserAssetsWithoutchain(address userAddresss) external view returns (Asset[] memory) {
        UserWallet storage userWallet = userWallets[userAddresss];
        Asset[] memory assets = new Asset[](userWallet.assetCountWithoutChain);
        for (uint256 i = 0; i < userWallet.assetCountWithoutChain; i++) {
            assets[i] = userWallet.assetsWithoutChain[i];
        }
        return assets;
    }

    function addNewAsset(address userAddress,uint128 sign,uint128 chainId) public {
        UserWallet storage userWallet = userWallets[userAddress];
        require(!isAssetExistByChain(userAddress,sign,chainId),"Asset is already exist");
        AssetPerChain memory assetPerChain;
        assetPerChain.sign = sign;
        assetPerChain.amount = 0;
        assetPerChain.chainId = chainId;
        userWallet.assets[userWallet.assetCount] = assetPerChain;
        userWallet.assetCount = userWallet.assetCount + 1;
        if(isAssetNotExist(userAddress,sign)){
            Asset memory asset;
            asset.sign = sign;
            asset.amount = 0;
            userWallet.assetsWithoutChain[userWallet.assetCountWithoutChain] = asset;
            userWallet.assetCountWithoutChain = userWallet.assetCountWithoutChain + 1;
        }
    }

    function depositAsset(address userAddress,uint128 sign,uint256 amount,uint128 chainId) public {
        UserWallet storage userWallet = userWallets[userAddress];
        for (uint256 i = 0; i < userWallet.assetCount; i++) {
            if(userWallet.assets[i].sign == sign && userWallet.assets[i].chainId == chainId){
                userWallet.assets[i].amount = userWallet.assets[i].amount + amount;
            }
        }
        for (uint256 i = 0; i < userWallet.assetCountWithoutChain; i++) {
            if(userWallet.assetsWithoutChain[i].sign == sign){
                userWallet.assetsWithoutChain[i].amount = userWallet.assetsWithoutChain[i].amount + amount;
            }
        }
    }

    function withdrawAsset(address userAddress,uint128 sign,uint256 amount,uint128 chainId) public {
        UserWallet storage userWallet = userWallets[userAddress];
        for (uint256 i = 0; i < userWallet.assetCount; i++) {
            if(userWallet.assets[i].sign == sign && userWallet.assets[i].chainId == chainId){
                userWallet.assets[i].amount = userWallet.assets[i].amount - amount;
            }
        }
        for (uint256 i = 0; i < userWallet.assetCountWithoutChain; i++) {
            if(userWallet.assetsWithoutChain[i].sign == sign){
                userWallet.assetsWithoutChain[i].amount = userWallet.assetsWithoutChain[i].amount - amount;
            }
        }
    }

    function isAssetExistByChain(address userAddress,uint128 sign,uint128 chainId) public  view returns(bool){
        UserWallet storage userWallet = userWallets[userAddress];
        for (uint256 i = 0; i < userWallet.assetCount; i++) {
            if(userWallet.assets[i].sign == sign && chainId == userWallet.assets[i].chainId){
                return true;
            }
        }
        return false;
    }


    function isAssetNotExist(address userAddress,uint128 sign) public  view returns(bool){
        UserWallet storage userWallet = userWallets[userAddress];
        for (uint256 i = 0; i < userWallet.assetCountWithoutChain; i++) {
            if(userWallet.assetsWithoutChain[i].sign == sign){
                return false;
            }
        }
        return true;
    }


}