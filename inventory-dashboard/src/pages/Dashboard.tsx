import {
    Grid,
    Card,
    CardContent,
    Typography
} from "@mui/material"


function Dashboard(){

    const cards = [

        {
            title:"Active Items",
            value:"0"
        },

        {
            title:"Current Inventory",
            value:"0"
        },

        {
            title:"Low Stock Alerts",
            value:"0"
        },

        {
            title:"Purchase Recommendations",
            value:"0"
        }

    ]

    return (
        <>
            <Typography
                variant="h4"
                gutterBottom
            >

                Dashboard

            </Typography>

            <Grid
                container
                spacing={3}
            >

                {
                    cards.map((card)=>(

                        <Grid
                            size={{
                                xs: 12,
                                sm: 6,
                                md: 3
                            }}
                            key={card.title}
                        >
                            <Card>
                                <CardContent>
                                    <Typography
                                        color="text.secondary"
                                    >

                                        {card.title}

                                    </Typography>

                                    <Typography
                                        variant="h4"
                                    >

                                        {card.value}

                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>
                    ))
                }
            </Grid>
        </>
    )
}

export default Dashboard