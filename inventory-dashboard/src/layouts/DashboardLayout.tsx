import {
    Box
} from "@mui/material"

import Sidebar from "../components/Sidebar"
import Header from "../components/Header"


function DashboardLayout({
    children
}: {
    children: React.ReactNode
}) {

    return (
        <Box
            sx={{
                display:"flex"
            }}
        >
            <Header />
            <Sidebar />
            <Box
                component="main"

                sx={{

                    flexGrow:1,

                    p:3,

                    mt:8

                }}

            >

                {children}

            </Box>
        </Box>
    )
}


export default DashboardLayout