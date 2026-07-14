import {
    useEffect,
    useState
} from "react"

import {
    Typography,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Paper
} from "@mui/material"

import {
    getItems,
    type Item
} from "../services/inventoryService"


function Inventory(){
    const [items,setItems] = useState<Item[]>([])

    useEffect(()=>{

        loadItems()

    },[])

    async function loadItems(){
        try{

            const data = await getItems()

            setItems(data)

        }
        catch(error){
            console.error(error)
        }
    }

    return (
        <>
            <Typography
                variant="h4"
                gutterBottom
            >
                Inventory
            </Typography>

            <TableContainer
                component={Paper}
            >
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>
                                Item Number
                            </TableCell>

                            <TableCell>
                                Description
                            </TableCell>

                            <TableCell>
                                Category
                            </TableCell>

                            <TableCell>
                                Active
                            </TableCell>
                        </TableRow>
                    </TableHead>

                    <TableBody>

                    {
                        items.map((item)=>(

                            <TableRow
                                key={item.id}
                            >

                                <TableCell>
                                    {item.itemNumber}
                                </TableCell>

                                <TableCell>
                                    {item.description}
                                </TableCell>

                                <TableCell>
                                    {item.category}
                                </TableCell>

                                <TableCell>
                                    {
                                        item.isActive
                                        ? "Yes"
                                        : "No"
                                    }
                                </TableCell>
                            </TableRow>
                        ))
                    }
                    </TableBody>
                </Table>
            </TableContainer>
        </>
    )
}


export default Inventory