import { RotateLoader } from "react-spinners";

const Loader = () => {
    return (
        <div
            style={{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                padding: "20px",
                height: "500px",
            }}
        >
            <RotateLoader color="black" size={30} />
        </div>
    );
};

export default Loader;
