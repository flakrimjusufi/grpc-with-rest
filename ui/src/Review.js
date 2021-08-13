import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import Grid from '@material-ui/core/Grid';

const Review= ({ reviewValues}) => {

  const personalInformation = [
    { name: 'First Name', desc: 'required', value: reviewValues['firstName'] },
    { name: 'Last Name', desc: 'optional', value: reviewValues['lastName'] },
    { name: 'Date Of Birth', desc: 'required', value: reviewValues['dateOfBirth'].toLocaleDateString() },
    { name: 'Phone Number', desc: 'required', value: reviewValues['phoneNumber'] },
    { name: 'Social Security Number', desc: 'required', value: reviewValues['socialSecurityNumber'] },
    { name: 'Employment Type', desc: 'required', value: reviewValues['employmentType'] },
    { name: 'Occupation', desc: 'optional', value: reviewValues['occupation'] },
    { name: 'Monthly Income', desc: 'required', value: reviewValues['monthlyIncome'] },
    { name: 'Years Employed', desc: 'required', value: reviewValues['yearsEmployed'] },
  ]

  const residencyInformation = [
    { name: 'Street Address', desc: 'optional', value: reviewValues['streetAddress'] },
    { name: 'Years at Address', desc: 'optional', value: reviewValues['yearsAtAddress'] },
    { name: 'City', desc: 'optional', value: reviewValues['city'] },
    { name: 'State', desc: 'optional', value: reviewValues['state'] },
    { name: 'Zip', desc: 'required', value: reviewValues['zip'] },
    { name: 'Country', desc: 'required', value: reviewValues['country'] },
    { name: 'Ownership', desc: 'required', value: reviewValues['ownership'] },
    { name: 'Monthly Payment', desc: 'required', value: reviewValues['monthlyPayment'] },
  ];

  const payments = [
    { name: 'Card Name', detail: reviewValues['cardName'] },
    { name: 'Card Type', detail: reviewValues['cardType'] },
    { name: 'Institution/Branch', detail: reviewValues['branch'] },
    { name: 'Credit Card Branding', detail: reviewValues['cardBranding'] },
  ];

  const useStyles = makeStyles((theme) => ({
    listItem: {
      padding: theme.spacing(1, 0),
    },
    title: {
      marginTop: theme.spacing(2),
    },
  }));


  const classes = useStyles();

  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom>
        Your personal information
      </Typography>
      <List disablePadding>
        {personalInformation.map((personalInfo) => (
          <ListItem className={classes.listItem} key={personalInfo.name}>
            <ListItemText primary={personalInfo.name} secondary={personalInfo.desc} />
            <Typography variant="body2">{personalInfo.value}</Typography>
          </ListItem>
        ))}
      </List>
      <Typography variant="h6" gutterBottom>
        Your residency information
      </Typography>
      <List disablePadding>
        {residencyInformation.map((residencyInfo) => (
          <ListItem className={classes.listItem} key={residencyInfo.name}>
            <ListItemText primary={residencyInfo.name} secondary={residencyInfo.desc} />
            <Typography variant="body2">{residencyInfo.value}</Typography>
          </ListItem>
        ))}
      </List>
      <Grid container spacing={2}>
        <Grid item container direction="column" xs={12} sm={6}>
          <Typography variant="h6" gutterBottom className={classes.title}>
            Your Credit Card information
          </Typography>
          <Grid container>
            {payments.map((payment) => (
              <React.Fragment key={payment.name}>
                <Grid item xs={6}>
                  <Typography gutterBottom>{payment.name}</Typography>
                </Grid>
                <Grid item xs={6}>
                  <Typography gutterBottom>{payment.detail}</Typography>
                </Grid>
              </React.Fragment>
            ))}
          </Grid>
        </Grid>
      </Grid>
    </React.Fragment>
  );
}
export default Review;