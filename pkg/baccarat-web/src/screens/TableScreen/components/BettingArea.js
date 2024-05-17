import Grid from "@mui/material/Grid";
import Typography from "@mui/material/Typography";
import Chip from "./Chip";

const styles = {
  winContainerStyle: {
    height: 250,
    backgroundColor: "purple",
    margin: "20px",
  },
  itemContainerStyle: {
    height: 250,
    backgroundColor: "#CCCCCC",
    margin: "20px",
  },
  headerStyle: {
    display: "flex",
    flex: 1,
    justifyContent: "center",
  },
};

const WinContent = (props) => {
  const { item } = props;

  return (
    <Grid item xs={3.5} style={styles.winContainerStyle}>
      <div style={styles.headerStyle}>
        <div>
          <Typography variant="h2" component="h2">
            ({item.odd})
          </Typography>
          <Typography variant="h2" component="h2">
            {item.text}
          </Typography>
        </div>
      </div>
      <Chip value={100} />
    </Grid>
  );
};

const NormalContent = (props) => {
  const { item } = props;

  return (
    <Grid item xs={3.5} style={styles.itemContainerStyle}>
      <div style={styles.headerStyle}>
        <div>
          <Typography variant="h2" component="h2">
            ({item.odd})
          </Typography>
          <Typography variant="h2" component="h2">
            {item.text}
          </Typography>
        </div>
      </div>

      <Chip value={100} />
    </Grid>
  );
};

const BettingArea = (props) => {
  return (
    <Grid container spacing={1} justifyContent="center">
      {props.items.map((item, index) => {
        if (index === 1)
          return <WinContent key={`betting-content-${index}`} item={item} />;
        return <NormalContent key={`betting-content-${index}`} item={item} />;
      })}
    </Grid>
  );
};

export default BettingArea;
