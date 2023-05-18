// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

contract LendingPool {

    struct Deposit {
        uint256 timestamp;
        uint256 balance;
        uint256 duration;
    }
    struct Repayment {
        uint256 amount;
        uint256 dueDate;
        bool paid;
    }
    uint256 totalSupply = 1000;
    uint256 interestRate = 12;
    mapping(address => Deposit[]) public deposits;
    mapping(address => uint256) public balances;
    mapping(address => uint256) public borrowed;
    mapping(address => uint256) public totalBarrowedLimit;
    mapping(address => Repayment[]) public repayments;

    uint256 public collateralizationRatio;
    uint256 public totalDeposits;

    event Borrow(address borrower, uint256 amount, uint256 interestRate, uint256 loanTerm, uint256 totalRepaymentAmount, uint256 dueDate);

    constructor(uint256 _collateralizationRatio) {
        collateralizationRatio = _collateralizationRatio;
    }

    function decreaseTotalBarrowedLimitOf(address user,uint256 amount) public{
        require(totalBarrowedLimit[user] - amount > 0,"Total limit cannot be less than zero");
        totalBarrowedLimit[user] -= amount;
    }

    function increaseTotalBarrowedLimitOf(address user,uint256 amount) public{
        totalBarrowedLimit[user] += amount;
    }

    function getInterestRate() view external returns(uint256){
        return interestRate;
    }

    function setInterestRate(uint256 _interestRate) public{
        interestRate = _interestRate;
    }

    //   MONEY TRANSFER 
    function deposit() public payable {
        balances[msg.sender] += msg.value;
        totalDeposits += msg.value;
    }

    function withdraw(uint256 amount) public {
        require(balances[msg.sender] >= amount, "Insufficient balance");
        balances[msg.sender] -= amount;
        payable(msg.sender).transfer(amount);
    }


    // LENDING
    function deposit(uint256 amount) public payable{
        require(amount > 0, "Value must be bigger than 0");
        deposits[msg.sender].push(Deposit(block.timestamp, amount, 0));
        totalSupply += amount;
    }

    function withdraw() public {
        Deposit[] memory depositAmount = deposits[msg.sender];
        uint256 totalAmount = 0;
        for (uint j = depositAmount.length; j != 0; j -= 10) {  //for loop example
            uint256 duration = block.timestamp - depositAmount[j].timestamp;
            totalAmount += depositAmount[j].balance + interestRate;
            depositAmount[j].balance = 0;
            depositAmount[j].duration = duration;
        }
        require(totalSupply > totalAmount, "Total supply is not enough");
        payable(msg.sender).transfer(totalAmount);
    }

    // BARROW
    function borrow(uint256 amount, uint256 loanTerm) public {
        // calculate the minimum required collateral based on the collateralization ratio
        uint256 requiredCollateral = amount * collateralizationRatio;

        // check that the user has sufficient collateral
        require(balances[msg.sender] >= requiredCollateral, "Insufficient collateral");

        require(totalBarrowedLimit[msg.sender] > amount,"Insufficient barrow limit");

         // check that there is enough deposit available to borrow
        require(totalDeposits >= amount, "Insufficient deposit");

        // check that the user is not already borrowing the maximum allowed amount
        //require(borrowed[msg.sender] + amount <= borrowingLimit[msg.sender], "Exceeded borrowing limit");

        // calculate the total amount to be repaid, including interest
        uint256 totalRepaymentAmount = amount + (amount * interestRate * loanTerm) / 100;

         // subtract the borrowed amount from totalDeposits
        totalDeposits -= amount;

        // transfer the borrowed amount to the user's address
        payable(msg.sender).transfer(amount);

        // update the user's borrowed amount and repayment schedule
        borrowed[msg.sender] += amount;
        repayments[msg.sender].push(Repayment(totalRepaymentAmount, block.timestamp + loanTerm, false));

        // emit an event to log the borrow transaction
        emit Borrow(msg.sender, amount, interestRate, loanTerm, totalRepaymentAmount, block.timestamp + loanTerm);
    }

    function calculateInterest(uint256 principal, uint256 duration) public view returns (uint256) {
        uint256 decimalInterestRate = interestRate / 100;
        uint256 interest = (principal * decimalInterestRate * duration) / 365 days;
        return interest;
    }

    function repay(uint256 index) public payable {
        require(index < repayments[msg.sender].length, "Invalid repayment index");
        Repayment storage repayment = repayments[msg.sender][index];
        require(!repayment.paid, "Repayment already paid");
        require(msg.value == repayment.amount, "Incorrect repayment amount");
        repayment.paid = true;
    }

    
}
