import { NavLink } from "react-router-dom";

import { BsLayoutSidebarInset } from "react-icons/bs";
import { RxDashboard } from "react-icons/rx";
import { IoSearchOutline } from "react-icons/io5";
import { FaRegChartBar } from "react-icons/fa";
import { CgDigitalocean } from "react-icons/cg";

import "./Menu.scss";
import { useState } from "react";

const Menu = () => {
    const [isOpen, setIsOpen] = useState(false);

    const handleOpenMenu = () => {
        setIsOpen(!isOpen);
    };

    const handleNavClick = () => {
        if (!isOpen) setIsOpen(true);
    };

    return (
        <div className={`menu-container ${isOpen ? "open" : "closed"}`}>
            <div className="menu-header">
                <a href="/" className="logo">
                    <CgDigitalocean color="var(--yellow)" size={30} />
                    {isOpen && <h1>G-Scores</h1>}
                </a>
                {isOpen && (
                    <div className="sidebar-icon" onClick={handleOpenMenu}>
                        <BsLayoutSidebarInset color="white" size={18} />
                    </div>
                )}
            </div>
            <nav>
                <NavLink
                    to="/"
                    onDoubleClick={handleNavClick}
                    className={({ isActive }) => (isActive ? "active" : "")}
                >
                    <RxDashboard color="white" size={18} />
                    {isOpen && <span>Dashboard</span>}
                </NavLink>

                <NavLink
                    to="/search-page"
                    onDoubleClick={handleNavClick}
                    className={({ isActive }) => (isActive ? "active" : "")}
                >
                    <IoSearchOutline color="white" size={18} />
                    {isOpen && <span>Search Scores</span>}
                </NavLink>

                <NavLink
                    to="/report"
                    onDoubleClick={handleNavClick}
                    className={({ isActive }) => (isActive ? "active" : "")}
                >
                    <FaRegChartBar color="white" size={18} />
                    {isOpen && <span>Reports</span>}
                </NavLink>
            </nav>
        </div>
    );
};

export default Menu;
