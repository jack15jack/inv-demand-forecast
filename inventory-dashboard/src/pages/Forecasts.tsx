import {
    useState
} from "react"

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
    LineChart,
    Line,
    XAxis,
    YAxis,
    CartesianGrid,
    Tooltip,
    Legend,
    ResponsiveContainer
} from "recharts"

import {
    getForecast,
    type ForecastResponse
} from "../services/forecastService"

function Forecasts(){

    const [itemID,setItemID] = useState(1)

    const [forecast,setForecast] = useState<ForecastResponse | null>(null)

    async function loadForecast(){

        try{

            const data =
                await getForecast(
                    itemID,
                    365,
                    7
                )

            console.log(data)

            setForecast(data)
        }
        catch(error){
            console.error(error)
        }
    }

    function buildChartData(){

        if(!forecast)
            return []

        const historyLength = forecast.historicalDemand.length

        const recentHistory = forecast.historicalDemand.slice(historyLength - 14)

        const historical =recentHistory.map(
                (value,index)=>({
                    day:`-${14-index}`,
                    historical:value,
                    forecast:null
                })
            )


        const future =forecast.dailyForecast.map(
                (value,index)=>({
                    day:`+${index+1}`,
                    historical:null,
                    forecast:value
                })
            )

        return [
            ...historical,
            ...future
        ]
    }

    return (
        <>
            <Typography
                variant="h4"
                gutterBottom
            >
                Forecasts
            </Typography>

            <Stack
                spacing={2}
                direction="row"
                sx={{
                    mb:3
                }}

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
                    onClick={loadForecast}
                >
                    Generate Forecast
                </Button>

            </Stack>
            {
                forecast &&
                <>
                    <Grid
                        container
                        spacing={2}
                    >
                        <Grid
                            size={{
                                xs:12,
                                md:3
                            }}
                        >
                            <Card>
                                <CardContent>
                                    <Typography>
                                        Current Stock
                                    </Typography>

                                    <Typography
                                        variant="h4"
                                    >
                                        {forecast.currentStock}
                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>

                        <Grid
                            size={{
                                xs:12,
                                md:3
                            }}
                        >
                            <Card>
                                <CardContent>
                                    <Typography>
                                        7 Day Demand Forecast
                                    </Typography>

                                    <Typography
                                        variant="h4"
                                    >
                                        {forecast.forecastedDemand}
                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>

                        <Grid
                            size={{
                                xs:12,
                                md:3
                            }}
                        >
                            <Card>
                                <CardContent>
                                    <Typography>
                                        Ending Inventory
                                    </Typography>

                                    <Typography
                                        variant="h4"
                                    >
                                        {forecast.predictedEndingInventory}
                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>

                        <Grid
                            size={{
                                xs:12,
                                md:3
                            }}
                        >
                            <Card>
                                <CardContent>
                                    <Typography>
                                        Daily Demand
                                    </Typography>

                                    <Typography
                                        variant="h4"
                                    >
                                        {
                                            forecast.dailyDemand.toFixed(2)
                                        }
                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>

                        <Grid
                            size={{
                                xs:12,
                                md:3
                            }}
                        >
                            <Card>
                                <CardContent>
                                    <Typography>
                                        Forecast Confidence
                                    </Typography>

                                    <Typography
                                        variant="h4"
                                    >
                                        {
                                            forecast.confidence.score
                                        }%
                                    </Typography>

                                    <Typography>
                                        {
                                            forecast.confidence.level
                                        }
                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>
                    </Grid>

                    <Card
                        sx={{
                            mt:3,
                            height:450
                        }}
                    >
                        <CardContent
                            sx={{
                                height:"100%"
                            }}
                        >
                            <ResponsiveContainer
                                width="100%"
                                height="100%"
                            >
                                <LineChart
                                    data={
                                        buildChartData()
                                    }
                                >
                                    <CartesianGrid />

                                    <XAxis
                                        dataKey="day"
                                    />

                                    <YAxis/>

                                    <Tooltip />

                                    <Legend />

                                    <Line
                                        type="monotone"
                                        dataKey="historical"
                                        name="Historical Demand"
                                    />

                                    <Line
                                        type="monotone"
                                        dataKey="forecast"
                                        name="Forecast Demand"
                                    />
                                </LineChart>
                            </ResponsiveContainer>
                        </CardContent>
                    </Card>
                </>
            }
        </>
    )
}

export default Forecasts