import { Link } from "react-router-dom";

export default function Nav() {
	return(
		<>	
        	<div className="ml-5 col-lg-1 col-md-3 float-start"><Link to="/" >Dashboard</Link></div>
        	<div className="ml-5 col-lg-1 col-md-3 float-start"><Link to="/post" >Add Posts</Link></div>
    	</>
	);
}