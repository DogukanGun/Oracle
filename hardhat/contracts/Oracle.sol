// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

contract Oracle {

    address lendingPoolContract;

    function setInterestRate(uint interestRate) external returns(bool){
        bytes memory payload = abi.encodeWithSignature("setInterestRate(uint256)",interestRate);
        (bool success,) = lendingPoolContract.call(payload);
        require(success);
        return success;
    }

    function increaseTotalBarrowedLimitOf(address user, uint interestRate) external returns(bool){
        bytes memory payload = abi.encodeWithSignature("increaseTotalBarrowedLimitOf(address,uint256)",user,interestRate);
        (bool success,) = lendingPoolContract.call(payload);
        require(success);
        return success;
    }

    function decreaseTotalBarrowedLimitOf(address user, uint interestRate) external returns(bool){
        bytes memory payload = abi.encodeWithSignature("decreaseTotalBarrowedLimitOf(address,uint256)",user,interestRate);
        (bool success,) = lendingPoolContract.call(payload);
        require(success);
        return success;
    }

    function setLendingPoolAddress(address _lendingPoolAddress) external {
        lendingPoolContract = _lendingPoolAddress;
    }

    function getLendingPoolAddress() external view returns(address){
        return lendingPoolContract;
    }
}