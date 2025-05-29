import { Routes, Route } from "react-router-dom";
import Header from "../Header/Header";
import Dashboard from "../../pages/DashboardPage/Dashboard";
import Search from "../../pages/Search/Search";
import "./MainContainer.scss";
import Report from "../../pages/Report/Report";

const MainContainer = () => {
    return (
        <div className="main-container">
            <Header />

            <div className="content">
                <Routes>
                    <Route path="/" element={<Dashboard />} />
                    <Route path="/search-page" element={<Search />} />
                    <Route path="/report" element={<Report />} />
                </Routes>
            </div>
        </div>
    );
};

export default MainContainer;
