import {
    Typography,
    Card,
    CardContent,
    Grid
} from "@mui/material"


function Home(){

    return (

        <>

            <Typography
                variant="h4"
                gutterBottom
            >
                Inventory Forecasting System
            </Typography>


            <Typography
                variant="body1"
                sx={{mb:4}}
            >
                Monitor inventory levels, analyze demand,
                and make data-driven purchasing decisions.
            </Typography>



            <Grid
                container
                spacing={2}
            >

                <Grid size={{xs:12, md:4}}>

                    <Card>

                        <CardContent>

                            <Typography>
                                Forecasting Engine
                            </Typography>

                            <Typography variant="h6">
                                Holt Linear Smoothing
                            </Typography>

                        </CardContent>

                    </Card>

                </Grid>


                <Grid size={{xs:12, md:4}}>

                    <Card>

                        <CardContent>

                            <Typography>
                                Database
                            </Typography>

                            <Typography variant="h6">
                                PostgreSQL
                            </Typography>
                        </CardContent>
                    </Card>
                </Grid>

                <Grid size={{xs:12, md:4}}>
                    <Card>
                        <CardContent>
                            <Typography>
                                System Status
                            </Typography>

                            <Typography variant="h6">
                                Online
                            </Typography>
                        </CardContent>
                    </Card>
                </Grid>
            </Grid>
        </>
    )
}


export default Home