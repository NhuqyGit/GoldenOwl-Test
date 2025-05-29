import { useState } from "react";
import { API_HOST } from "../../common/config";
import "./search.scss";

const Search = () => {
    const [sbd, setSbd] = useState("");
    const [loading, setLoading] = useState(false);
    const [result, setResult] = useState(null);

    const handleSearch = async () => {
        if (!sbd) return;

        setLoading(true);
        setResult(null);

        try {
            console.log(API_HOST);
            const response = await fetch(`/student-scores?sbd=${sbd}`, {
                method: "GET",
            });

            const data = await response.json();
            console.log(data);
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
                {loading && <p>Loading...</p>}

                {result && !result.error && (
                    <div className="result-table">
                        <p>
                            <strong>SBD:</strong> {result.SBD}
                        </p>
                        <p>
                            <strong>Toán:</strong> {result.Toan ?? "N/A"}
                        </p>
                        <p>
                            <strong>Ngữ Văn:</strong> {result.NguVan ?? "N/A"}
                        </p>
                        <p>
                            <strong>Ngoại Ngữ:</strong>{" "}
                            {result.NgoaiNgu ?? "N/A"} ({result.MaNgoaiNgu})
                        </p>
                        <p>
                            <strong>Lịch Sử:</strong> {result.LichSu ?? "N/A"}
                        </p>
                        <p>
                            <strong>Địa Lí:</strong> {result.DiaLi ?? "N/A"}
                        </p>
                        <p>
                            <strong>GDCD:</strong> {result.GDCD ?? "N/A"}
                        </p>
                        <p>
                            <strong>Vật Lí:</strong> {result.VatLi ?? "N/A"}
                        </p>
                        <p>
                            <strong>Hóa Học:</strong> {result.HoaHoc ?? "N/A"}
                        </p>
                        <p>
                            <strong>Sinh Học:</strong> {result.SinhHoc ?? "N/A"}
                        </p>
                    </div>
                )}

                {result?.error && <p className="error">{result.error}</p>}
            </div>
        </div>
    );
};

export default Search;
