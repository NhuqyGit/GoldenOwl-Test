import React, { createContext, useContext, useState, useEffect } from "react";
import { API_HOST } from "../common/config";

const StudentScoreContext = createContext();

export const useStudentScores = () => useContext(StudentScoreContext);

export const StudentScoreProvider = ({ children }) => {
    const [reportData, setReportData] = useState(null);
    const [top10GroupA, setTop10GroupA] = useState([]);

    const [loadingReport, setLoadingReport] = useState(true);
    const [loadingTop10, setLoadingTop10] = useState(true);

    useEffect(() => {
        // fetch api report
        const fetchReport = async () => {
            setLoadingReport(true);
            try {
                const res = await fetch(`${API_HOST}/student-scores/report`);
                const json = await res.json();
                const subjectMap = {
                    toan: "Math",
                    ngu_van: "Literature",
                    ngoai_ngu: "Foreign Language",
                    vat_li: "Physics",
                    hoa_hoc: "Chemistry",
                    sinh_hoc: "Biology",
                    lich_su: "History",
                    dia_li: "Geography",
                    gdcd: "Civic Education",
                };
                const transformed = Object.entries(json).map(([key, val]) => ({
                    subject: subjectMap[key] || key,
                    ...val,
                }));
                setReportData(transformed);
            } catch (err) {
                console.error("Fetch report error:", err);
            } finally {
                setLoadingReport(false);
            }
        };

        //fetch top10
        const fetchTop10 = async () => {
            setLoadingTop10(true);
            try {
                const res = await fetch(
                    `${API_HOST}/student-scores/top-10-groupA`
                );
                const json = await res.json();
                setTop10GroupA(json);
            } catch (err) {
                console.error("Fetch top10 error:", err);
            } finally {
                setLoadingTop10(false);
            }
        };

        fetchReport();
        fetchTop10();
    }, []);

    return (
        <StudentScoreContext.Provider
            value={{
                reportData,
                top10GroupA,
                loadingReport,
                loadingTop10,
            }}
        >
            {children}
        </StudentScoreContext.Provider>
    );
};
