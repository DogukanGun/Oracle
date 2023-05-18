import React from "react";
import { Outlet } from 'react-router-dom';
import CustomNavbar from "../navbar/CustomNavbar";

const CustomContent = () => {

    return(
        <React.Fragment>
            <CustomNavbar/>
            <main>
                <Outlet />
            </main>
        </React.Fragment>
    )

}

export default CustomContent;