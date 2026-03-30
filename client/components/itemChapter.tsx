import { RSC_HEADER } from "next/dist/client/components/app-router-headers";
import * as React from 'react';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import Avatar from '@mui/material/Avatar';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import Chip from '@mui/material/Chip';
import Link from '@mui/material/Link';

import { ItemChapter } from "./listChapter";
export default function ComponetItemChapter({ item }: { item: ItemChapter }) {
  // Generate a default badge color if status is active/completed
  const isCompleted = item.Status?.toLowerCase() === 'completed';
  const statusColor = isCompleted ? 'success' : 'primary';
  const href_item =  `/chapter/detail/${String(item.ID)}`

  return (
    <Link href= {href_item} underline="none">
      <ListItem
        alignItems="flex-start"
        sx={{
          borderRadius: 2,
          transition: 'all 0.2s ease-in-out',
          '&:hover': {
            bgcolor: 'action.hover',
            transform: 'translateX(4px)'
          },
          flexDirection: { xs: 'column', sm: 'row' },
          py: 2
        }}
      >
        <ListItemAvatar sx={{ mt: { xs: 0, sm: 1 }, mb: { xs: 2, sm: 0 } }}>
          <Avatar
            alt={item.Title || "Chapter"}
            src={`https://ui-avatars.com/api/?name=${item.Title || 'C'}&background=random`}
            sx={{ width: 56, height: 56, mr: 2 }}
          />
        </ListItemAvatar>
        <ListItemText
          primary={
            <Box display="flex" justifyContent="space-between" alignItems="center" flexWrap="wrap" mb={1}>
              <Typography variant="h6" fontWeight="bold" color="text.primary">
                {item.Title || "Untitled Chapter"}
              </Typography>
              {item.Status && (
                <Chip
                  label={item.Status}
                  color={statusColor as "default" | "success" | "primary"}
                  size="small"
                  sx={{ fontWeight: 'bold', ml: { xs: 0, sm: 2 }, mt: { xs: 1, sm: 0 } }}
                />
              )}
            </Box>
          }
          secondary={
            <React.Fragment>
              <Typography
                component="span"
                variant="body2"
                sx={{ color: 'text.secondary', display: 'block', lineHeight: 1.6 }}
              >
                {item.Body ? item.Body : "No description available for this chapter."}
              </Typography>
            </React.Fragment>
          }
          sx={{ width: '100%', m: 0 }}
        />
      </ListItem>
    </Link>
  );
}