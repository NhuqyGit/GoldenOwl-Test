import React, { createContext, useContext, useState, useEffect } from "react";

const StudentScoreContext = createContext();

export const useStudentScores = () => useContext(StudentScoreContext);

export const StudentScoreProvider = ({ children }) => {
    const [reportData, setReportData] = useState(null);
    const [top10GroupA, setTop10GroupA] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchAll = async () => {
            setLoading(true);
            try {
                // Fetch api report
                const reportResponse = await fetch("/student-scores/report");
                const reportJson = await reportResponse.json();

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

                const transformed = Object.entries(reportJson).map(
                    ([subjectKey, levels]) => ({
                        subject: subjectMap[subjectKey] || subjectKey,
                        ...levels,
                    })
                );
                setReportData(transformed);

                // Fetch top 10 group A
                const top10Response = await fetch(
                    "/student-scores/top-10-groupA"
                );
                const top10Json = await top10Response.json();
                setTop10GroupA(top10Json);
            } catch (err) {
                console.error("Fetch error:", err);
            } finally {
                setLoading(false);
            }
        };

        fetchAll();
    }, []);

    return (
        <StudentScoreContext.Provider
            value={{
                reportData,
                top10GroupA,
                loading,
            }}
        >
            {children}
        </StudentScoreContext.Provider>
    );
};
