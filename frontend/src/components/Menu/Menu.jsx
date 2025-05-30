import { NavLink } from "react-router-dom";

import { BsLayoutSidebarInset } from "react-icons/bs";
import { RxDashboard } from "react-icons/rx";
import { IoSearchOutline } from "react-icons/io5";
import { FaRegChartBar } from "react-icons/fa";
import { CgDigitalocean } from "react-icons/cg";
import { HiOutlineMenuAlt3 } from "react-icons/hi";

import "./Menu.scss";
import { useState, useEffect } from "react";

const Menu = () => {
    const [isOpen, setIsOpen] = useState(false);
    const [windowWidth, setWindowWidth] = useState(window.innerWidth);

    const isT1 = windowWidth > 780;
    const isT2 = windowWidth <= 580 && windowWidth > 430;
    const isT3 = windowWidth <= 430;

    useEffect(() => {
        const handleResize = () => {
            setWindowWidth(window.innerWidth);
        };
        window.addEventListener("resize", handleResize);
        return () => window.removeEventListener("resize", handleResize);
    }, []);

    const handleOpenMenu = () => {
        if (isT3) {
            setIsOpen(!isOpen); // dọc
        } else if (isT1) {
            setIsOpen(!isOpen); // ngang
        }
        // T2 không cho mở
    };

    const handleNavClick = () => {
        if (isT3) {
            setIsOpen(false); // đóng menu sau khi chọn item ở T3
        }
    };

    const handNavDoubleClick = () => {
        if (isT1 && !isOpen) {
            setIsOpen(true); // ở T1 thì mở nếu đang đóng
        }
    };

    return (
        <>
            {isT3 && !isOpen && (
                <div className="mobile-menu-button" onClick={handleOpenMenu}>
                    <HiOutlineMenuAlt3 size={28} color="var(--black)" />
                </div>
            )}

            <div
                className={`menu-container 
                    ${isOpen ? "open" : "closed"} 
                    ${isT3 ? "mobile" : ""} `}
            >
                <div className="menu-header">
                    <a href="/" className="logo">
                        <CgDigitalocean color="var(--yellow)" size={30} />
                        {(isOpen || isT3) && <h1>G-Scores</h1>}
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
                        onDoubleClick={handNavDoubleClick}
                        onClick={handleNavClick}
                        className={({ isActive }) => (isActive ? "active" : "")}
                    >
                        <RxDashboard color="white" size={18} />
                        {(isOpen || isT3) && <span>Dashboard</span>}
                    </NavLink>

                    <NavLink
                        to="/search-page"
                        onDoubleClick={handNavDoubleClick}
                        onClick={handleNavClick}
                        className={({ isActive }) => (isActive ? "active" : "")}
                    >
                        <IoSearchOutline color="white" size={18} />
                        {(isOpen || isT3) && <span>Search Scores</span>}
                    </NavLink>

                    <NavLink
                        to="/report"
                        onDoubleClick={handNavDoubleClick}
                        onClick={handleNavClick}
                        className={({ isActive }) => (isActive ? "active" : "")}
                    >
                        <FaRegChartBar color="white" size={18} />
                        {(isOpen || isT3) && <span>Reports</span>}
                    </NavLink>
                </nav>
            </div>
        </>
    );
};

export default Menu;
