import {
    Toolbar,
    Drawer,
    List,
    ListItem,
    ListItemButton,
    ListItemText
} from "@mui/material"

import { Link } from "react-router-dom"


const menuItems = [
    {
        name: "Dashboard",
        path: "/"
    },
    {
        name: "Transactions",
        path: "/transactions"
    },
    {
        name: "Inventory",
        path: "/inventory"
    },

    {
        name: "Analytics",
        path: "/analytics"
    },
    
    {
        name: "Forecasts",
        path: "/forecasts"
    },
    {
        name: "Recommendations",
        path: "/recommendations"
    }    
]


function Sidebar(){
    return (

        <Drawer
            variant="permanent"
            sx={{
                width:240,
                "& .MuiDrawer-paper":{
                    width:240,
                    boxSizing:"border-box"
                }
            }}
        >
            <Toolbar />
            <List>
                {
                    menuItems.map((item)=>(
                        <ListItem
                            key={item.path}
                            disablePadding
                        >
                            <ListItemButton
                                component={Link}
                                to={item.path}
                            >

                                <ListItemText
                                    primary={item.name}
                                />
                            </ListItemButton>
                        </ListItem>
                    ))
                }
            </List>
        </Drawer>
    )
}


export default Sidebar