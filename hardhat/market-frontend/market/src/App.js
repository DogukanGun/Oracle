import logo from './logo.svg';
import './App.css';
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import LendingPoolComponent from './feature/lending_pool/LendingPoolComponent';
import CustomContent from './layout/content/CustomContent';
import { Children } from 'react';
import BarrowMoney from './feature/barrow/BarrowMoneyComponent';

const router = createBrowserRouter([
  {
    path: "/",
    element:<CustomContent/>,
    children: [
      {
        index:true,
        element: <LendingPoolComponent/>
      },
      {
        path: "barrow",
        element: <BarrowMoney/>
      }
    ]
  }
]);
function App() {
  return (
    <RouterProvider router={router}/>
  );
}

export default App;
