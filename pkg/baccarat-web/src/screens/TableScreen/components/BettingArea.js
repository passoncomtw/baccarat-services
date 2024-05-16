import Grid from "@mui/material/Grid";
import Typography from "@mui/material/Typography";
import Chip from "./Chip";

const styles = {
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

const BettingArea = (props) => {
  return (
    <Grid container spacing={1} justifyContent="center">
      {props.items.map((item) => (
        <Grid item xs={3.5} style={styles.itemContainerStyle}>
          <div
            style={styles.headerStyle}
          >
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
      ))}
    </Grid>
  );
};

export default BettingArea;
