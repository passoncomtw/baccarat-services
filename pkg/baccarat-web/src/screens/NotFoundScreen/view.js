const NotFoundScreen = () => (
  <div style={{textAlign: "center"}}>
    <img
      src="/assets/page-wrong.png"
      alt=""
      style={{width: "250px", margin: "auto", paddingTop: "3em"}}
    />
    <h2 style={{fontSize: "1.8rem", color: "#515151"}}>伺服器維修中</h2>
    <p style={{fontSize: "1em", color: "#515151"}}>
      嗨，你會看到這個畫面是因為我們的伺服器正在維修中。
      <br />
      我們會盡快讓他恢復工作，造成您的不便，敬請見諒，謝謝！
    </p>
    <p style={{fontSize: "1em", color: "#515151"}}>
      Sorry! We are maintaining the server. It'll come back online soon.
    </p>
  </div>
);

export default NotFoundScreen;
