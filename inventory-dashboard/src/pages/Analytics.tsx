import { useState } from "react"

import {
    Typography,
    TextField,
    Button,
    Card,
    CardContent,
    Grid,
    Stack
} from "@mui/material"

import {
    getAnalytics,
    type AnalyticsResponse
} from "../services/analyticsService"

function Analytics() {
    const [itemID, setItemID] = useState(1)

    const [analytics, setAnalytics] = useState<AnalyticsResponse | null>(null)

    async function loadAnalytics() {

        try {
            const data = await getAnalytics(
                itemID,
                365
            )

            setAnalytics(data)
        }
        catch(err){
            console.error(err)
        }
    }

    return (
        <>
            <Typography
                variant="h4"
                gutterBottom
            >
                Analytics
            </Typography>

            <Stack
                direction="row"
                spacing={2}
                sx={{ mb: 3 }}
            >
                <TextField
                    label="Item ID"
                    type="number"
                    value={itemID}
                    onChange={(e)=>
                        setItemID(
                            Number(e.target.value)
                        )
                    }
                />

                <Button
                    variant="contained"
                    onClick={loadAnalytics}
                >
                    Load Analytics
                </Button>
            </Stack>
            {
                analytics &&
                <>
                    <Grid
                        container
                        spacing={2}
                    >
                        <Grid size={{ xs:12, md:4 }}>
                            <Card>
                                <CardContent>
                                    <Typography>
                                        Current Stock
                                    </Typography>

                                    <Typography variant="h4">
                                        {analytics.currentStock}
                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>

                        <Grid size={{ xs:12, md:4 }}>
                            <Card>
                                <CardContent>
                                    <Typography>
                                        Units Sold
                                    </Typography>

                                    <Typography variant="h4">
                                        {analytics.unitsSold}
                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>

                        <Grid size={{ xs:12, md:4 }}>
                            <Card>
                                <CardContent>
                                    <Typography>
                                        Analysis Window
                                    </Typography>

                                    <Typography variant="h4">
                                        {analytics.analysisWindowDays} Days
                                    </Typography>

                                </CardContent>
                            </Card>
                        </Grid>

                        <Grid size={{ xs:12, md:4 }}>
                            <Card>
                                <CardContent>
                                    <Typography>
                                        Average Daily Demand
                                    </Typography>

                                    <Typography variant="h4">
                                        {analytics.averageDailyDemand.toFixed(2)}
                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>

                        <Grid size={{ xs:12, md:4 }}>
                            <Card>
                                <CardContent>
                                    <Typography>
                                        Average Weekly Demand
                                    </Typography>

                                    <Typography variant="h4">
                                        {analytics.averageWeeklyDemand.toFixed(2)}
                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>

                        <Grid size={{ xs:12, md:4 }}>
                            <Card>
                                <CardContent>
                                    <Typography>
                                        Days Remaining
                                    </Typography>

                                    <Typography variant="h4">
                                        {analytics.daysOfInventoryRemaining.toFixed(1)}
                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>

                        <Grid size={{ xs:12 }}>
                            <Card>
                                <CardContent>
                                    <Typography
                                        variant="h6"
                                        gutterBottom
                                    >
                                        Last Sale
                                    </Typography>

                                    <Typography variant="h5">
                                        {

                                            analytics.lastSale
                                            ?

                                            new Date(
                                                analytics.lastSale
                                            ).toLocaleString()

                                            :

                                            "No sales recorded"
                                        }

                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>
                    </Grid>
                </>
            }
        </>
    )
}

export default Analytics