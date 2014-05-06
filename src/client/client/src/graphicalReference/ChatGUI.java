package graphicalReference;

import java.awt.BorderLayout;
import java.awt.Color;
import java.awt.Dimension;
import java.awt.event.ActionEvent;

import javax.swing.JFrame;
import javax.swing.JOptionPane;
import javax.swing.JPanel;
import javax.swing.JScrollPane;
import javax.swing.JTextField;
import javax.swing.JTextPane;
import javax.swing.border.EmptyBorder;

import org.json.JSONException;

import clientNetworking.NetManager;

import packageManaging.Message;
import packageManaging.MessageHandler;

public class ChatGUI extends GUI {
	public JTextPane chatPanel = new JTextPane();
	public JTextField field;
	public String userName;
	public NetManager NM;
	@Override
	public void render() {
		
	}
	
	public void render(JPanel j, NetManager NM, JFrame JF){
		this.NM = NM;
		chatStart(j, NM, JF);
	}

	@Override
	public void actionPerformed(ActionEvent e) {
		String str = this.field.getText();
		
		MessageHandler l = (MessageHandler)this.NM.handler;
		Message mess = new Message(str, this.userName, 10);
		String sender = "";
		try {
			sender = l.encode(mess);
		} catch (JSONException e1) {
			e1.printStackTrace();
		}
		
		this.NM.send(sender);
		this.field.setText("");
		
		
	}
	public void chatStart(JPanel j, NetManager NM, JFrame JF){
		
		
       
		j.setLayout(new BorderLayout());
        this.field = new JTextField(20);
        field.addActionListener(this);
        field.setBackground(Color.DARK_GRAY.brighter());
        field.setBorder(new EmptyBorder(2,0,10,0));
        JScrollPane js = new JScrollPane(chatPanel);
        Dimension preferredSize = new Dimension(400, 200);
        js.setPreferredSize(preferredSize);
        js.setAutoscrolls(true);
        chatPanel.setBackground(Color.DARK_GRAY);
        j.setBorder(null);
    ;
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
	
	
	
}
