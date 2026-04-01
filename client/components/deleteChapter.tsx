"use client";
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import { ItemChapter } from "./listChapter";
import DeleteIcon from '@mui/icons-material/Delete';
import * as React from 'react';
import Box from '@mui/material/Box';
export default function DeleteChapter({ item }: { item: ItemChapter}) {
  const [openDeleteDialog, setOpenDeleteDialog] = React.useState(false);
  const handleDeleteConfirm = async () => {
    const params_object = {
      'id': String(item.ID),
    };
    const queryString = new URLSearchParams(params_object).toString();
    var url = `/api/chapter/delete?${queryString}`;
    const response = await fetch(url ,  {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      },
    })
  };
  const handleDeleteClick = () => {
    setOpenDeleteDialog(true);
  };

  const handleDeleteClose = () => {
    setOpenDeleteDialog(false);
  };
    return <>
    <Box sx={{
        display: 'flex',
        justifyContent: 'center', // Horizontal centering
        alignItems: 'center',     // Vertical centering
      }}>
        <IconButton onClick={handleDeleteClick} color="error" aria-label="delete">
          <DeleteIcon />
        </IconButton>
    </Box>
    <Dialog
        open={openDeleteDialog}
        onClose={handleDeleteClose}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">
          {"Delete Chapter?"}
        </DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
            Are you sure you want to delete the chapter "{item.Title}"? This action cannot be undone.
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleDeleteClose}>Cancel</Button>
          <Button onClick={handleDeleteConfirm} color="error" autoFocus>
            Delete
          </Button>
        </DialogActions>
      </Dialog>
    </>
     
}