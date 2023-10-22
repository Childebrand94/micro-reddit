import FilterOptions from "./FilterOptions";
import NavProfile from "./NavProfile";
import NavSearch from "./NavSearch";

const NavBar = () => {
  return (
    <nav className="w-full h-9 bg-gray-100 flex ">
      <div>
        <a href="/">
          <img src="../assets/logo-reddit.svg" />
        </a>
      </div>
      <FilterOptions />

      <div className="flex">
        <NavSearch />
        <NavProfile />
      </div>
    </nav>
  );
};
export default NavBar;
