import React from "react";
import "./Top10.scss";

const Top10 = ({ data }) => {
    if (!data || data.length === 0) {
        return <p>No data available.</p>;
    }

    return (
        <div className="top10-container">
            <table>
                <thead>
                    <tr>
                        <th>Rank</th>
                        <th>SBD</th>
                        <th>Math</th>
                        <th>Literature</th>
                        <th>English</th>
                        <th>Physics</th>
                        <th>Chemistry</th>
                        <th>History</th>
                        <th>Geography</th>
                        <th>Civic Education</th>
                        <th>Sum</th>
                    </tr>
                </thead>
                <tbody>
                    {data.map((student, index) => {
                        const total =
                            (student.Toan || 0) +
                            (student.VatLi || 0) +
                            (student.HoaHoc || 0);

                        return (
                            <tr key={student.ID}>
                                <td>{index + 1}</td>
                                <td>{student.SBD}</td>
                                <td>{student.Toan || "N/A"}</td>
                                <td>{student.NguVan || "N/A"}</td>
                                <td>{student.NgoaiNgu || "N/A"}</td>
                                <td>{student.VatLi || "N/A"}</td>
                                <td>{student.HoaHoc || "N/A"}</td>
                                <td>{student.LichSu || "N/A"}</td>
                                <td>{student.DiaLi || "N/A"}</td>
                                <td>{student.GDCD || "N/A"}</td>
                                <td>{total.toFixed(2)}</td>
                            </tr>
                        );
                    })}
                </tbody>
            </table>
        </div>
    );
};

export default Top10;
