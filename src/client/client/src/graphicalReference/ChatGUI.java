package graphicalReference;

import java.awt.BorderLayout;
import java.awt.Color;
import java.awt.Dimension;

import java.text.SimpleDateFormat;
import java.util.Calendar;

import javax.swing.JFrame;
import javax.swing.JOptionPane;
import javax.swing.JPanel;
import javax.swing.JScrollPane;
import javax.swing.JTextField;
import javax.swing.JTextPane;
import javax.swing.border.EmptyBorder;




public class ChatGUI{
	public JTextPane chatPanel = new JTextPane();
	public JTextField field;
	public String userName;
	

	public void render(JPanel j, JFrame JF){
		
		chatStart(j, JF);
	}


	public void chatStart(JPanel j,JFrame JF){
		
		
       
		j.setLayout(new BorderLayout());
        this.field = new JTextField(20);
        
        field.setBackground(Color.DARK_GRAY.brighter());
        field.setBorder(new EmptyBorder(2,0,10,0));
        field.setActionCommand("chatmessage");
        JScrollPane js = new JScrollPane(chatPanel);
        Dimension preferredSize = new Dimension(400, 200);
        js.setPreferredSize(preferredSize);
        js.setAutoscrolls(true);
        chatPanel.setBackground(Color.DARK_GRAY);
        j.setBorder(null);
    
        js.getViewport().setBorder(null);
        chatPanel.setBorder(null);
        field.setBorder(null);
        field.setEditable(true);
        j.setVisible(true);
        chatPanel.setEditable(false);
        chatPanel.setForeground(Color.LIGHT_GRAY);

        
        j.add(js, "Center");
        j.add(field, "South");
        JF.add(j );
        JF.validate();
        userName = (String)JOptionPane.showInputDialog("Write username!");
        
		
        
		

	}
	
	public void chatUpdate(String mess,String username){
				
    	Calendar cal = Calendar.getInstance();
    	cal.getTime();
    	SimpleDateFormat sdf = new SimpleDateFormat("HH:mm:ss");
    	
		String prevMess = chatPanel.getText();
		chatPanel.setText(prevMess + "<" + sdf.format(cal.getTime())+ ">" + username + ":" + mess + "\n");
		chatPanel.setCaretPosition(chatPanel.getDocument().getLength());
	}
	
	
	
}