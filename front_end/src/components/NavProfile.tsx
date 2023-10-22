const NavProfile = () => {
  const loggedIn: boolean = true;
  return (
    <div>
      {loggedIn ? (
        <div className="bg-gray-400 p-1 rounded-xl text-xl">LogIn</div>
      ) : (
        <image path="../assets/logo-reddit.svg" />
      )}
    </div>
  );
};
export default NavProfile;
