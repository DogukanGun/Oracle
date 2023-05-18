import React, { useState } from "react"
import { ethers } from "ethers"
import LendingPool from "./LendingPool.json"

const LENDING_POOL_ADDRESS = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
const LendingPoolComponent = () => {
    const [walletAddress, setWalletAddress] = useState("");
    const [amount, setAmountAddress] = useState("");
    async function lendMoney() {
        await connectToContract();
    }

    async function connectToContract() {
        if (typeof window.ethereum !== "undefined") {
            //ethereum is usable get reference to the contract
            const provider = new ethers.providers.Web3Provider(window.ethereum);
            const contract = new ethers.Contract(LENDING_POOL_ADDRESS, LendingPool.abi, provider);
            try {
                const data = await contract.address;
                console.log("Data: ", data);
            } catch (e) {
                console.log("Err: ", e)
            }
        }
    }

    // @ts-ignore
    return (
        <>
            <div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
                <div class="sm:mx-auto sm:w-full sm:max-w-sm">
                    <h2 class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">Send Money</h2>
                </div>

                <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
                    <form class="space-y-6" action="#" method="POST">
                        <div>
                            <label for="email" class="block text-sm font-medium leading-6 text-gray-900">Wallet Address</label>
                            <div class="mt-2">
                                <input type="text" class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" name="wallet_address" value={walletAddress} onChange={event => setWalletAddress(event.target.value)} />
                            </div>
                        </div>

                        <div>
                            <div class="flex items-center justify-between">
                                <label for="password" class="block text-sm font-medium leading-6 text-gray-900">Amount</label>
                            </div>
                            <div class="mt-2">
                                <input type="text" class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" name="amount" value={amount} onChange={(event => setAmountAddress(event.target.value))} />
                            </div>
                        </div>

                        <div>
                            <button class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600" onClick={lendMoney}>Lend Money</button>
                        </div>
                    </form>

                   
                </div>
            </div>
        </>
    )
}

export default LendingPoolComponent;