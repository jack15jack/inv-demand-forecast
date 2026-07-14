import {
    Typography,
    TextField,
    Button,
    MenuItem,
    Card,
    CardContent,
    Stack,
    Grid
} from "@mui/material"

import { useState } from "react"

import { createTransaction } from "../services/transactionService"
import { createItem } from "../services/itemService"


function Transactions(){
    const [itemForm,setItemForm] = useState({
        itemNumber:"",
        description:"",
        category:"",
        unitCost:0,
        unitPrice:0,
        minimumStock:0,
        safetyStock:0,
        isActive:true
    })

    const handleItemChange = (field:string, value:any)=>{

        setItemForm({

            ...itemForm,

            [field]:value

        })
    }

    const handleCreateItem = async()=>{

        try{

            const item = await createItem(itemForm)

            alert(
                `Item created: ${item.itemNumber}`
            )

            setItemForm({

                itemNumber:"",
                description:"",
                category:"",
                unitCost:0,
                unitPrice:0,
                minimumStock:0,
                safetyStock:0,
                isActive:true

            })
        }
        catch(error){
            console.error(error)

            alert(
                "Failed to create item"
            )
        }
    }

    const [form,setForm] = useState({
        ItemID:0,
        TransactionType:"",
        Direction:"",
        Quantity:0,
        Reference:"",
        Notes:""
    })


    const handleChange = (field:string, value:any)=>{

        setForm({

            ...form,

            [field]:value

        })
    }


    const handleSubmit = async()=>{

        try{

            await createTransaction(form)

            alert(
                "Transaction created"
            )

            setForm({
                ItemID:0,
                TransactionType:"",
                Direction:"",
                Quantity:0,
                Reference:"",
                Notes:""
            })
        }
        catch(error){

            console.error(error)

            alert(
                "Failed to create transaction"
            )
        }
    }

    return (
        <>
            <Typography
                variant="h4"
                gutterBottom
            >
                Transactions
            </Typography>

            <Card
                sx={{
                    mb:3
                }}
            >
                <CardContent>
                    <Typography
                        variant="h6"
                        gutterBottom
                    >
                        Create Item
                    </Typography>

                    <Grid
                        container
                        spacing={2}
                    >
                        <Grid size={{xs:12, md:6}}>
                            <TextField
                                fullWidth
                                label="Item Number"
                                value={itemForm.itemNumber}
                                onChange={(e)=>
                                    handleItemChange(
                                        "itemNumber",
                                        e.target.value
                                    )
                                }
                            />
                        </Grid>

                        <Grid size={{xs:12, md:6}}>
                            <TextField
                                fullWidth
                                label="Description"
                                value={itemForm.description}
                                onChange={(e)=>
                                    handleItemChange(
                                        "description",
                                        e.target.value
                                    )
                                }
                            />
                        </Grid>

                        <Grid size={{xs:12, md:6}}>
                            <TextField
                                fullWidth
                                label="Category"
                                value={itemForm.category}
                                onChange={(e)=>
                                    handleItemChange(
                                        "category",
                                        e.target.value
                                    )
                                }
                            />
                        </Grid>

                        <Grid size={{xs:12, md:3}}>
                            <TextField
                                fullWidth
                                label="Unit Cost"
                                type="number"
                                value={itemForm.unitCost}
                                onChange={(e)=>
                                    handleItemChange(
                                        "unitCost",
                                        Number(e.target.value)
                                    )
                                }
                            />
                        </Grid>

                        <Grid size={{xs:12, md:3}}>
                            <TextField
                                fullWidth
                                label="Unit Price"
                                type="number"
                                value={itemForm.unitPrice}
                                onChange={(e)=>
                                    handleItemChange(
                                        "unitPrice",
                                        Number(e.target.value)
                                    )
                                }
                            />
                        </Grid>

                        <Grid size={{xs:12, md:3}}>
                            <TextField
                                fullWidth
                                label="Minimum Stock"
                                type="number"
                                value={itemForm.minimumStock}
                                onChange={(e)=>
                                    handleItemChange(
                                        "minimumStock",
                                        Number(e.target.value)
                                    )
                                }
                            />
                        </Grid>

                        <Grid size={{xs:12, md:3}}>
                            <TextField
                                fullWidth
                                label="Safety Stock"
                                type="number"
                                value={itemForm.safetyStock}
                                onChange={(e)=>
                                    handleItemChange(
                                        "safetyStock",
                                        Number(e.target.value)
                                    )
                                }
                            />
                        </Grid>
                    </Grid>

                    <Button
                        sx={{mt:3}}
                        variant="contained"
                        onClick={handleCreateItem}
                    >
                        Create Item
                    </Button>
                </CardContent>
            </Card>

            <Card>
                <CardContent>
                    <Typography
                        variant="h6"
                        gutterBottom
                    >
                        Log Transaction
                    </Typography>

                    <Stack
                        spacing={2}
                    >
                        <TextField
                            label="Item ID"
                            type="number"
                            value={form.ItemID}
                            onChange={(e)=>
                                handleChange(
                                    "ItemID",
                                    Number(e.target.value)
                                )
                            }
                        />

                        <TextField
                            select
                            label="Transaction Type"
                            value={form.TransactionType}
                            onChange={(e)=>
                                handleChange(
                                    "TransactionType",
                                    e.target.value
                                )
                            }
                        >
                            <MenuItem value="SALE">
                                Sale
                            </MenuItem>

                            <MenuItem value="PURCHASE">
                                Purchase
                            </MenuItem>

                            <MenuItem value="RETURN">
                                Return
                            </MenuItem>

                            <MenuItem value="ADJUSTMENT">
                                Adjustment
                            </MenuItem>

                            <MenuItem value="TRANSFER">
                                Transfer
                            </MenuItem>
                        </TextField>

                        <TextField
                            select
                            label="Direction"
                            value={form.Direction}
                            onChange={(e)=>
                                handleChange(
                                    "Direction",
                                    e.target.value
                                )
                            }
                        >
                            <MenuItem value="IN">
                                Inbound
                            </MenuItem>

                            <MenuItem value="OUT">
                                Outbound
                            </MenuItem>
                        </TextField>

                        <TextField
                            label="Quantity"
                            type="number"
                            value={form.Quantity}
                            onChange={(e)=>
                                handleChange(
                                    "Quantity",
                                    Number(e.target.value)
                                )
                            }
                        />

                        <TextField
                            label="Reference"
                            value={form.Reference}
                            onChange={(e)=>
                                handleChange(
                                    "Reference",
                                    e.target.value
                                )
                            }
                        />

                        <TextField
                            label="Notes"
                            multiline
                            rows={3}
                            value={form.Notes}
                            onChange={(e)=>
                                handleChange(
                                    "Notes",
                                    e.target.value
                                )
                            }
                        />

                        <Button
                            variant="contained"
                            onClick={handleSubmit}
                        >
                            Log Transaction
                        </Button>
                    </Stack>
                </CardContent>
            </Card>
        </>
    )
}


export default Transactions