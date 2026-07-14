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
    Paper,
    Chip
} from "@mui/material"

import {
    getRecommendations,
    type PurchaseRecommendation
} from "../services/recommendationService"

function Recommendations(){

    const [
        recommendations,
        setRecommendations
    ] = useState<PurchaseRecommendation[]>([])

    useEffect(()=>{
        loadRecommendations()
    },[])

    async function loadRecommendations(){

        try{

            const data = await getRecommendations()
            setRecommendations(data)

        }
        catch(error){
            console.error(error)
        }
    }

    function urgencyColor(
        urgency:string
    ){
        switch(urgency){
            case "HIGH":
                return "error"

            case "MEDIUM":
                return "warning"

            default:
                return "success"
        }
    }

    return (
        <>
            <Typography
                variant="h4"
                gutterBottom
            >
                Purchase Recommendations
            </Typography>

            <TableContainer
                component={Paper}
            >
                <Table
                    size="small"
                >
                    <TableHead>
                        <TableRow>
                            <TableCell>
                                Item
                            </TableCell>

                            <TableCell>
                                Description
                            </TableCell>

                            <TableCell>
                                Stock
                            </TableCell>

                            <TableCell>
                                30 Day Forecast
                            </TableCell>

                            <TableCell>
                                Projected
                            </TableCell>

                            <TableCell>
                                Safety Stock
                            </TableCell>

                            <TableCell>
                                Purchase
                            </TableCell>

                            <TableCell>
                                Urgency
                            </TableCell>

                            <TableCell>
                                Reason
                            </TableCell>
                        </TableRow>
                    </TableHead>

                    <TableBody>
                    {
                        recommendations.map((rec)=>(
                            <TableRow
                                key={rec.itemID}
                            >
                                <TableCell>
                                    {rec.itemNumber}
                                </TableCell>

                                <TableCell>
                                    {rec.description}
                                </TableCell>

                                <TableCell>
                                    {rec.currentStock}
                                </TableCell>

                                <TableCell>
                                    {rec.forecastedDemand}
                                    {" units"}
                                </TableCell>

                                <TableCell>
                                    {rec.projectedInventory}
                                </TableCell>

                                <TableCell>
                                    {rec.safetyStock}
                                </TableCell>

                                <TableCell>
                                    <strong>
                                        {rec.recommendedPurchase}
                                    </strong>
                                </TableCell>

                                <TableCell>
                                    <Chip
                                        label={
                                            rec.urgency
                                        }
                                        color={
                                            urgencyColor(
                                                rec.urgency
                                            )
                                        }
                                        size="small"
                                    />
                                </TableCell>

                                <TableCell>
                                    {rec.reason}
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

export default Recommendations