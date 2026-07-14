import {
    Typography,
    TextField,
    Button,
    MenuItem,
    Card,
    CardContent,
    Stack
} from "@mui/material"
import { useState } from "react"
import { createTransaction } from "../services/transactionService"

function Transactions(){

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

    const handleSubmit = async ()=>{
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