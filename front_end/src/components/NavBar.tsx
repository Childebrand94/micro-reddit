import FilterOptions from "./FilterOptions";
import NavProfile from "./NavProfile";
import NavSearch from "./NavSearch";

const NavBar = () => {
  return (
    <nav className="max-w-full h-9 bg-blue-100 flex ">
      <img
        className="m-1"
        src="../../public/assets/logo-reddit.svg"
        alt="Reddit Logo"
      />
      <FilterOptions />

      <div className="flex items-end">
        <NavSearch />
        <NavProfile />
      </div>
    </nav>
  );
};
export default NavBar;
