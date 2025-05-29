import "./Header.scss";
import { useLocation } from "react-router-dom";

const Header = () => {
    const location = useLocation();

    const getTitle = () => {
        switch (location.pathname) {
            case "/":
                return "Dashboard";
            case "/search-page":
                return "Search";
            case "/report":
                return "Reports";
        }
    };

    return (
        <div className="header-container">
            <h1>{getTitle()}</h1>
        </div>
    );
};

export default Header;
