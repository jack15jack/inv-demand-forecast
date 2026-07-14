import {
    BrowserRouter,
    Routes,
    Route
} from "react-router-dom"

import DashboardLayout from "./layouts/DashboardLayout"

import Dashboard from "./pages/Dashboard"
import Inventory from "./pages/Inventory"
import Transactions from "./pages/Transactions"
import Forecasts from "./pages/Forecasts"
import Recommendations from "./pages/Recommendations"
import Analytics from "./pages/Analytics"

function App(){
    return (
        <BrowserRouter>
            <DashboardLayout>
                <Routes>
                    <Route
                        path="/"
                        element={<Dashboard />}
                    />
                    <Route
                        path="/transactions"
                        element={<Transactions />}
                    />
                    <Route
                        path="/inventory"
                        element={<Inventory />}
                    />
                    
                    <Route
                        path="/forecasts"
                        element={<Forecasts />}
                    />
                    <Route
                        path="/recommendations"
                        element={<Recommendations />}
                    />
                    <Route
                        path="/analytics"
                        element={<Analytics />}
                    />
                </Routes>
            </DashboardLayout>
        </BrowserRouter>
    )
}

export default App