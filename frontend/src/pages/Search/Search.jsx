import { useState } from "react";
import { API_HOST } from "../../common/config";
import { ClipLoader } from "react-spinners";
import "./Search.scss";

const Search = () => {
    const [sbd, setSbd] = useState("");
    const [loading, setLoading] = useState(false);
    const [result, setResult] = useState(null);

    const handleSearch = async () => {
        if (!sbd) return;

        setLoading(true);
        setResult(null);

        try {
            const response = await fetch(`/student-scores?sbd=${sbd}`);
            const data = await response.json();
            setResult(data);
        } catch (error) {
            console.error("Fetch error:", error);
            setResult({ error: "Failed to fetch data." });
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="search-container">
            <div className="search-bar">
                <h2>User Registration</h2>
                <label>Registration Number</label>
                <div className="input">
                    <input
                        type="text"
                        placeholder="Enter registration number"
                        value={sbd}
                        onChange={(e) => setSbd(e.target.value)}
                    />
                    <button onClick={handleSearch}>Submit</button>
                </div>
            </div>

            <div className="search-result">
                <h2>Search Result</h2>

                {loading && (
                    <div
                        style={{
                            display: "flex",
                            justifyContent: "center",
                            padding: "20px",
                        }}
                    >
                        <ClipLoader color="#36d7b7" size={40} />
                    </div>
                )}

                {result && !result.error && (
                    <div className="result-table">
                        <p>
                            <strong>SBD:</strong> {result.SBD}
                        </p>
                        <p>
                            <strong>Math:</strong> {result.Toan ?? "N/A"}
                        </p>
                        <p>
                            <strong>Literature:</strong>{" "}
                            {result.NguVan ?? "N/A"}
                        </p>
                        <p>
                            <strong>English:</strong> {result.NgoaiNgu ?? "N/A"}{" "}
                            ({result.MaNgoaiNgu})
                        </p>
                        <p>
                            <strong>Physics:</strong> {result.VatLi ?? "N/A"}
                        </p>
                        <p>
                            <strong>Chemistry:</strong> {result.HoaHoc ?? "N/A"}
                        </p>
                        <p>
                            <strong>Biology:</strong> {result.SinhHoc ?? "N/A"}
                        </p>
                        <p>
                            <strong>History:</strong> {result.LichSu ?? "N/A"}
                        </p>
                        <p>
                            <strong>Geography:</strong> {result.DiaLi ?? "N/A"}
                        </p>
                        <p>
                            <strong>Civic Education:</strong>{" "}
                            {result.GDCD ?? "N/A"}
                        </p>
                    </div>
                )}

                {result?.error && <p className="error">{result.error}</p>}
            </div>
        </div>
    );
};

export default Search;
